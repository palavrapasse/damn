package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/palavrapasse/damn/pkg/entity"
)

const (
	_sqliteDriverName = "sqlite3"
)

type DatabaseContext[R Record] struct {
	DB       *sql.DB
	FilePath string
}

type TransactionContext[R Record] struct {
	Tx *sql.Tx
}

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

func (ctx DatabaseContext[R]) NewTransactionContext() (TransactionContext[R], error) {
	tx, err := ctx.DB.Begin()

	return TransactionContext[R]{Tx: tx}, err
}

func Convert[R Record, T User](ctx DatabaseContext[R]) DatabaseContext[T] {
	return DatabaseContext[T]{}
}

func (ctx DatabaseContext[Record]) Insert(i Import) error {

	// ctxd := Convert(ctx)

	// ctxd.CustomQuery("", func() (User, []any) { return User{}, []any{} })

	var tx *sql.Tx

	tctx, err := ctx.NewTransactionContext()

	tx = tctx.Tx

	defer func() {
		if err != nil {
			log.Printf("could not complete transaction: %v", err)

			err = tx.Rollback()
		}

		if err != nil {
			log.Printf("could not rollback transaction: %v", err)
		}
	}()

	func() {
		us := make([]User, len(i.AffectedUsers))
		cr := make([]Credentials, len(i.AffectedUsers))

		j := 0

		for u, c := range i.AffectedUsers {
			us[j] = u
			cr[j] = c

			j++
		}

		// Primary first

		ut := NewUserTable(us)
		ct := NewCredentialsTable(cr)
		bat := NewBadActorTable(i.Leakers)
		lt := NewLeakTable(i.Leak)
		pt := NewPlatformTable(i.AffectedPlatforms)

		ptt := []any{ut, ct, bat, lt, pt}

		for j, t := range ptt {
			tc := t.(PrimaryTable[Record])
			t, err = tctx.insertPrimary(tc)

			if err == nil {
				t, err = tctx.findPrimary(tc)
			}

			if err == nil {
				ptt[j] = t
			} else {
				return
			}
		}

		ut = ptt[0].(PrimaryTable[User])
		ct = ptt[1].(PrimaryTable[Credentials])
		bat = ptt[2].(PrimaryTable[BadActor])
		lt = ptt[3].(PrimaryTable[Leak])
		pt = ptt[4].(PrimaryTable[Platform])

		// Foreign now

		us = ut.Records
		cr = ct.Records
		bas := bat.Records
		ls := lt.Records
		ps := pt.Records

		l := ls[0]
		afu := map[User]Credentials{}

		for k := range us {
			afu[us[k]] = cr[k]
		}

		hct := NewHashCredentialsTable(cr)
		hut := NewHashUserTable(us)
		lbat := NewLeakBadActorTable(map[Leak][]BadActor{l: bas})
		lcrt := NewLeakCredentialsTable(map[Leak][]Credentials{l: cr})
		lptt := NewLeakPlatformTable(map[Leak][]Platform{l: ps})
		lut := NewLeakUserTable(map[Leak][]User{l: us})
		uct := NewUserCredentialsTable(afu)

		ftt := []any{hct, hut, lbat, lcrt, lptt, lut, uct}

		for j, t := range ftt {
			tc := t.(ForeignTable[Record])
			t, err = tctx.insertForeign(tc)

			if err == nil {
				ftt[j] = t
			} else {
				return
			}
		}
	}()

	if err == nil {
		err = tx.Commit()
	}

	return err
}

func (ctx DatabaseContext[R]) CustomQuery(q string, mp func() (R, []any), v ...any) ([]R, error) {
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
			err = fmt.Errorf("could not complete transaction: %w", err)

			err = tx.Rollback()
		}

		if err != nil {
			err = fmt.Errorf("could not rollback transaction: %w", err)
		}
	}()

	if err == nil {
		func() {
			stmt, err = tx.Prepare(q)

			if err != nil {
				return
			}

			rs, err = stmt.Query(v)

			if err != nil {
				return
			}

			for rs.Next() {
				r, addrs := mp()

				err = rs.Scan(addrs...)

				if err != nil {
					break
				} else {
					rcs = append(rcs, r)
				}
			}
		}()
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
