package database

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/palavrapasse/damn/pkg/entity"
	. "github.com/palavrapasse/damn/pkg/entity/query"
	. "github.com/palavrapasse/damn/pkg/entity/subscribe"
)

const (
	_sqliteDriverName = "sqlite3"
)

const (
	errorMessageCompleteTransaction = "could not complete transaction: %w"
	errorMessageRollbackTransaction = "could not rollback transaction: %w"
)

const MaxElementsOfGoroutine = 5000

type DatabaseContext[R Record] struct {
	DB       *sql.DB
	FilePath string
}

type TransactionContext[R Record] struct {
	Tx *sql.Tx
}

type TypedQueryResultMapper[R Record] func() (*R, []any)
type AnonymousErrorCallback func() (any, error)

func NewDatabaseContext[R Record](fp string) (DatabaseContext[R], error) {
	db, err := sql.Open(_sqliteDriverName, fp)

	if err == nil {
		err = db.Ping()
	}

	return DatabaseContext[R]{
		DB:       db,
		FilePath: fp,
	}, err
}

func Convert[R Record, T Record](ctx DatabaseContext[R]) DatabaseContext[T] {
	return DatabaseContext[T](ctx)
}

func (ctx DatabaseContext[R]) NewTransactionContext() (TransactionContext[R], error) {
	tx, err := ctx.DB.Begin()

	return TransactionContext[R]{Tx: tx}, err
}

func (ctx DatabaseContext[Record]) Insert(i Import) (AutoGenKey, error) {

	var tx *sql.Tx

	tctx, err := ctx.NewTransactionContext()

	tx = tctx.Tx

	defer func() {
		if err != nil {
			err = fmt.Errorf(errorMessageCompleteTransaction, err)

			err = tx.Rollback()
		}

		if err != nil {
			err = fmt.Errorf(errorMessageRollbackTransaction, err)
		}
	}()

	var leakId AutoGenKey

	func() {
		us := i.AffectedUsers

		// Primary first

		var pts []any

		cbs := []AnonymousErrorCallback{
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[User](tctx), NewConcurrentPrimaryTable(MaxElementsOfGoroutine, us, NewUserTable))
			},
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[BadActor](tctx), NewConcurrentPrimaryTable(MaxElementsOfGoroutine, i.Leakers, NewBadActorTable))
			},
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[Leak](tctx), NewLeakTable(i.Leak))
			},
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[Platform](tctx), NewConcurrentPrimaryTable(MaxElementsOfGoroutine, i.AffectedPlatforms, NewPlatformTable))
			},
		}

		pts, err = returnOnCallbackError(cbs)

		if err != nil {
			return
		}

		// Foreign now

		us = pts[0].(PrimaryTable[User]).Records
		bas := pts[1].(PrimaryTable[BadActor]).Records
		ls := pts[2].(PrimaryTable[Leak]).Records
		ps := pts[3].(PrimaryTable[Platform]).Records

		l := ls[0]
		leakId = l.LeakId

		cbs = []AnonymousErrorCallback{
			func() (any, error) {
				return typedInsertForeign(TransactionContext[HashUser](tctx), NewConcurrentHashForeignTable(MaxElementsOfGoroutine, us, NewHashUserTable))
			},
			func() (any, error) {
				return typedInsertForeign(TransactionContext[LeakBadActor](tctx), NewLeakBadActorTable(map[Leak][]BadActor{l: bas}))
			},
			func() (any, error) {
				return typedInsertForeign(TransactionContext[LeakPlatform](tctx), NewLeakPlatformTable(map[Leak][]Platform{l: ps}))
			},
			func() (any, error) {
				return typedInsertForeign(TransactionContext[LeakUser](tctx), NewLeakUserTable(map[Leak][]User{l: us}))
			},
		}

		_, err = returnOnCallbackError(cbs)

		if err != nil {
			return
		}
	}()

	if err == nil {
		err = tx.Commit()
	}

	return leakId, err
}

func (ctx DatabaseContext[Record]) InsertSubscription(s Subscription) error {

	var tx *sql.Tx

	tctx, err := ctx.NewTransactionContext()

	tx = tctx.Tx

	defer func() {
		if err != nil {
			err = fmt.Errorf(errorMessageCompleteTransaction, err)

			err = tx.Rollback()
		}

		if err != nil {
			err = fmt.Errorf(errorMessageRollbackTransaction, err)
		}
	}()

	func() {
		sub := s.Subscriber
		aff := s.Affected

		// Primary first

		var pts []any

		cbs := []AnonymousErrorCallback{
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[Subscriber](tctx), NewSubscriberTable(sub))
			},
			func() (any, error) {
				return typedInsertAndFindPrimary(TransactionContext[Affected](tctx), NewAffectedTable(aff))
			},
		}

		pts, err = returnOnCallbackError(cbs)

		if err != nil {
			return
		}

		// Foreign now

		sub = pts[0].(PrimaryTable[Subscriber]).Records[0]
		aff = pts[1].(PrimaryTable[Affected]).Records

		if len(aff) > 0 {
			cbs = []AnonymousErrorCallback{
				func() (any, error) {
					return typedInsertForeign(TransactionContext[SubscriberAffected](tctx), NewSubscriberAffectedTable(sub, aff))
				},
			}

			_, err = returnOnCallbackError(cbs)

			if err != nil {
				return
			}
		}
	}()

	if err == nil {
		err = tx.Commit()
	}

	return err
}

// Execute a query that can be customized using prepared statements. Consumers must provide a typed callback
// that shall return each row result mapped as a pointer to a struct.
//
// An example of a call would be:
//
//	ctx.CustomQuery("SELECT * FROM User WHERE email = ?", func() (*User, []any) {
//			u := User{}
//			return u, []any{&u.UserId, &u.Email}
//		}, email)
func (ctx DatabaseContext[R]) CustomQuery(q string, mp TypedQueryResultMapper[R], v ...any) ([]R, error) {
	var tctx TransactionContext[R]
	var tx *sql.Tx
	var rs *sql.Rows
	var stmt *sql.Stmt
	var err error
	var rcs []R

	tctx, err = ctx.NewTransactionContext()

	tx = tctx.Tx

	defer func() {
		if err != nil {
			err = fmt.Errorf(errorMessageCompleteTransaction, err)

			err = tx.Rollback()
		}

		if err != nil {
			err = fmt.Errorf(errorMessageRollbackTransaction, err)
		}
	}()

	if err == nil {
		func() {
			stmt, err = tx.Prepare(q)

			if err != nil {
				return
			}

			rs, err = stmt.Query(v...)

			if err != nil {
				return
			}

			for rs.Next() {
				r, addrs := mp()

				err = rs.Scan(addrs...)

				if err != nil {
					break
				}

				rcs = append(rcs, *r)
			}

			defer rs.Close()
		}()
	}

	if err == nil {
		err = tx.Commit()
	}

	return rcs, err
}

func (ctx TransactionContext[R]) findPrimary(t PrimaryTable[R]) (PrimaryTable[R], error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records[R]

	tx = ctx.Tx

	stmt, err = t.PrepareFindStatement(tx)

	if err == nil {
		records := t.Records

		for _, r := range records {
			if !t.HasPrimaryKeySet(r) {
				var row *sql.Row
				var rvp []*any
				var rvpp []any

				rv := t.Values(r)

				for _, v := range rv {
					var x = reflect.ValueOf(v).Interface()
					rvp = append(rvp, &x)
					rvpp = append(rvpp, &x)
				}

				row = stmt.QueryRow(t.FindValues(r)...)

				if row != nil {
					err = row.Scan(rvpp...)
				}

				if err == nil {
					rid, ok := (*rvp[0]).(int64)

					if ok {
						updatedRecords = append(updatedRecords, CopyWithNewKey(r, AutoGenKey(rid)))
					} else {
						err = fmt.Errorf("could not convert first query result to int64: %v", rvp)
					}
				} else {
					break
				}
			} else {
				updatedRecords = append(updatedRecords, r)
			}
		}
	}

	return t.Copy(updatedRecords), err
}

func (ctx TransactionContext[R]) insertPrimary(t PrimaryTable[R]) (PrimaryTable[R], error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records[R]

	tx = ctx.Tx

	stmt, err = t.PrepareInsertStatement(tx)

	if err == nil {
		records := t.Records

		for _, r := range records {
			var res sql.Result
			var raff int64
			var lid int64

			res, err = stmt.Exec(t.InsertValues(r)...)

			if res != nil {
				raff, err = res.RowsAffected()
			}

			if raff > 0 {
				lid, err = res.LastInsertId()
			}

			if err == nil {
				updatedRecords = append(updatedRecords, CopyWithNewKey(r, AutoGenKey(lid)))
			} else {
				break
			}
		}
	}

	return t.Copy(updatedRecords), err
}

func (ctx TransactionContext[R]) insertForeign(t ForeignTable[R]) (ForeignTable[R], error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records[R]

	tx = ctx.Tx

	stmt, err = t.PrepareInsertStatement(tx)

	if err == nil {
		records := t.Records

		for _, r := range records {
			_, err = stmt.Exec(t.InsertValues(r)...)

			if err != nil {
				break
			}
		}
	}

	return t.Copy(updatedRecords), err
}

func typedInsertAndFindPrimary[R BadActor | Leak | Platform | User | Subscriber | Affected](ctx TransactionContext[R], t PrimaryTable[R]) (PrimaryTable[R], error) {
	tctx := TransactionContext[R]{Tx: ctx.Tx}

	tu, err := tctx.insertPrimary(t)

	if err == nil {
		tu, err = tctx.findPrimary(tu)
	}

	return tu, err
}

func typedInsertForeign[R HashUser | LeakBadActor | LeakPlatform | LeakUser | SubscriberAffected](ctx TransactionContext[R], t ForeignTable[R]) (ForeignTable[R], error) {
	return ctx.insertForeign(t)
}

func returnOnCallbackError(cbs []AnonymousErrorCallback) ([]any, error) {
	var err error

	res := make([]any, len(cbs))

	for i, cb := range cbs {
		var t any

		t, err = cb()

		if err != nil {
			break
		}

		res[i] = t
	}

	return res, err
}
