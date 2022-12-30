package database

import (
	"database/sql"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/palavrapasse/damn/pkg/entity"
)

type DatabaseContext struct {
	DB       *sql.DB
	FilePath string
}

type TransactionContext struct {
	Tx *sql.Tx
}

const (
	_sqliteDriverName = "sqlite3"
)

func NewDatabaseContext(fp string) (DatabaseContext, error) {
	db, err := sql.Open(_sqliteDriverName, fp)

	if err == nil {
		err = db.Ping()
	}

	return DatabaseContext{
		DB:       db,
		FilePath: fp,
	}, err
}

func NewTransactionContext(db *sql.DB) (TransactionContext, error) {
	tx, err := db.Begin()

	return TransactionContext{Tx: tx}, err
}

func (ctx DatabaseContext) Insert(i Import) error {
	var tx *sql.Tx

	db := ctx.DB
	tctx, err := NewTransactionContext(db)

	tx = tctx.Tx

	defer func() {
		if err != nil {
			tx.Rollback()
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

		ptt := []PrimaryTable{ut, ct, bat, lt, pt}

		for j, t := range ptt {
			t, err = tctx.insertPrimary(t)

			if err == nil {
				t, err = tctx.findPrimary(t)
			}

			if err == nil {
				ptt[j] = t
			} else {
				return
			}
		}

		ut = ptt[0]
		ct = ptt[1]
		bat = ptt[2]
		lt = ptt[3]
		pt = ptt[4]

		// Foreign now

		us = ut.ToUserSlice()
		cr = ct.ToCredentialsSlice()
		bas := bat.ToBadActorSlice()
		ls := lt.ToLeakSlice()
		ps := pt.ToPlatformSlice()

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

		ftt := []ForeignTable{hct, hut, lbat, lcrt, lptt, lut, uct}

		for j, t := range ftt {
			t, err = tctx.insertForeign(t)

			if err == nil {
				ftt[j] = t
			} else {
				return
			}
		}
	}()

	if err == nil {
		tx.Commit()
	}

	return err
}

func (ctx TransactionContext) findPrimary(t PrimaryTable) (PrimaryTable, error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records

	tx = ctx.Tx

	stmt, err = t.PrepareFindStatement(tx)

	if err == nil {
		records := t.Records

		for _, r := range records {
			if !t.HasPrimaryKeySet(r) {
				var row *sql.Row
				var rid int64
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
					rid = (*rvp[0]).(int64)

					updatedRecords = append(updatedRecords, CopyWithNewKey(r, AutoGenKey(rid)))
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

func (ctx TransactionContext) insertPrimary(t PrimaryTable) (PrimaryTable, error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records

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

func (ctx TransactionContext) insertForeign(t ForeignTable) (ForeignTable, error) {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error

	var updatedRecords Records

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
