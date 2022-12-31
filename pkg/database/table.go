package database

import (
	"database/sql"
	"fmt"
	"strings"

	. "github.com/palavrapasse/damn/pkg/entity"
)

const (
	prepareInsertStatementSQLString         = "INSERT OR IGNORE INTO %s (%s) VALUES (%s)"
	prepareFindStatementSQLString           = "SELECT * FROM %s WHERE (%s) = (%s) LIMIT 1"
	prepareStatementPlaceholderSymbol       = "?"
	prepareStatementMultipleFieldsSeparator = ", "
)

const (
	unknownTableName = "void"
)

var unknownTableFields = []Field{}

type Table interface {
	Name() string
	Records() []Record
	Fields() []Field
	Values(Record) []any
	Copy(Records) DatabaseTable
	HasPrimaryKeySet(Record) bool
	InsertFields() []Field
	InsertValues(Record) []any
	FindFields() []Field
	FindValues(Record) []any
	PrepareInsertStatement(*sql.Tx) (*sql.Stmt, error)
	PrepareFindStatement(*sql.Tx) (*sql.Stmt, error)
}

type DatabaseTable struct {
	Table
	Records
}

type PrimaryTable DatabaseTable
type ForeignTable DatabaseTable

type BadActorTable = PrimaryTable
type CredentialsTable = PrimaryTable
type HashCredentialsTable = ForeignTable
type HashUserTable = ForeignTable
type LeakBadActorTable = ForeignTable
type LeakCredentialsTable = ForeignTable
type LeakPlatformTable = ForeignTable
type LeakTable = PrimaryTable
type LeakUserTable = ForeignTable
type PlatformTable = PrimaryTable
type UserCredentialsTable = ForeignTable
type UserTable = PrimaryTable

func NewBadActorTable(ba []BadActor) BadActorTable {
	rs := make(Records, len(ba))

	for i, v := range ba {
		rs[i] = v
	}

	return BadActorTable{Records: rs}
}

func NewCredentialsTable(cr []Credentials) CredentialsTable {
	rs := make(Records, len(cr))

	for i, v := range cr {
		rs[i] = v
	}

	return CredentialsTable{Records: rs}
}

func NewLeakTable(ls ...Leak) LeakTable {
	rs := make(Records, len(ls))

	for i, v := range ls {
		rs[i] = v
	}

	return LeakTable{Records: rs}
}

func NewPlatformTable(ps []Platform) PlatformTable {
	rs := make(Records, len(ps))

	for i, v := range ps {
		rs[i] = v
	}

	return PlatformTable{Records: rs}
}

func NewUserTable(us []User) UserTable {
	rs := make(Records, len(us))

	for i, v := range us {
		rs[i] = v
	}

	return UserTable{Records: rs}
}

func NewHashCredentialsTable(cr []Credentials) HashCredentialsTable {
	rs := make(Records, len(cr))

	for i, v := range cr {
		rs[i] = NewHashCredentials(v)
	}

	return HashCredentialsTable{Records: rs}
}

func NewHashUserTable(us []User) HashUserTable {
	rs := make(Records, len(us))

	for i, v := range us {
		rs[i] = NewHashUser(v)
	}

	return HashUserTable{Records: rs}
}

func NewLeakBadActorTable(lba map[Leak][]BadActor) LeakBadActorTable {
	rs := Records{}

	for l, bas := range lba {
		for _, ba := range bas {
			rs = append(rs, NewLeakBadActor(ba, l))
		}
	}

	return LeakBadActorTable{Records: rs}
}

func NewLeakCredentialsTable(lcr map[Leak][]Credentials) LeakCredentialsTable {
	rs := Records{}

	for l, crs := range lcr {
		for _, cr := range crs {
			rs = append(rs, NewLeakCredentials(cr, l))
		}
	}

	return LeakCredentialsTable{Records: rs}
}

func NewLeakPlatformTable(lpt map[Leak][]Platform) LeakPlatformTable {
	rs := Records{}

	for l, pts := range lpt {
		for _, pt := range pts {
			rs = append(rs, NewLeakPlatform(pt, l))
		}
	}

	return LeakPlatformTable{Records: rs}
}

func NewLeakUserTable(lus map[Leak][]User) LeakUserTable {
	rs := Records{}

	for l, us := range lus {
		for _, u := range us {
			rs = append(rs, NewLeakUser(u, l))
		}
	}

	return LeakUserTable{Records: rs}
}

func NewUserCredentialsTable(uc map[User]Credentials) UserCredentialsTable {
	rs := make(Records, len(uc))

	i := 0

	for u, c := range uc {
		rs[i] = UserCredentials{CredId: c.CredId, UserId: u.UserId}

		i++
	}

	return UserCredentialsTable{Records: rs}
}

func (bat BadActorTable) ToBadActorSlice() []BadActor {
	bas := make([]BadActor, len(bat.Records))

	for i, r := range bat.Records {
		ba := r.(BadActor)
		bas[i] = ba
	}

	return bas
}

func (ct CredentialsTable) ToCredentialsSlice() []Credentials {
	crs := make([]Credentials, len(ct.Records))

	for i, r := range ct.Records {
		cr := r.(Credentials)
		crs[i] = cr
	}

	return crs
}

func (lt LeakTable) ToLeakSlice() []Leak {
	ls := make([]Leak, len(lt.Records))

	for i, r := range lt.Records {
		l := r.(Leak)
		ls[i] = l
	}

	return ls
}

func (pt PlatformTable) ToPlatformSlice() []Platform {
	ps := make([]Platform, len(pt.Records))

	for i, r := range pt.Records {
		p := r.(Platform)
		ps[i] = p
	}

	return ps
}

func (ut UserTable) ToUserSlice() []User {
	us := make([]User, len(ut.Records))

	for i, r := range ut.Records {
		u := r.(User)
		us[i] = u
	}

	return us
}

func (pt PrimaryTable) Name() string {
	return DatabaseTable(pt).Name()
}

func (pt PrimaryTable) Fields() []Field {
	return DatabaseTable(pt).Fields()
}

func (pt PrimaryTable) Values(r Record) []any {
	return Values(r)
}

func (pt PrimaryTable) InsertFields() []Field {
	return DatabaseTable(pt).Fields()[1:]
}

func (pt PrimaryTable) InsertValues(r Record) []any {
	return Values(r)[1:]
}

func (pt PrimaryTable) FindFields() []Field {
	// todo: rely on sql tags
	return DatabaseTable(pt).Fields()[1:]
}

func (pt PrimaryTable) FindValues(r Record) []any {
	// todo: rely on sql tags
	return Values(r)[1:]
}

func (pt PrimaryTable) HasPrimaryKeySet(r Record) bool {
	// todo: rely on sql tags
	return Values(r)[0] != AutoGenKey(0)
}

func (pt PrimaryTable) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(pt.prepareInsertStatementString())

}

func (pt PrimaryTable) PrepareFindStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(pt.prepareFindStatementString())

}

func (pt PrimaryTable) Copy(rs Records) PrimaryTable {
	return PrimaryTable{Records: rs}
}

func (ft ForeignTable) Name() string {
	return DatabaseTable(ft).Name()
}

func (ft ForeignTable) Fields() []Field {
	return DatabaseTable(ft).Fields()
}

func (ft ForeignTable) InsertFields() []Field {
	return DatabaseTable(ft).Fields()
}

func (ft ForeignTable) InsertValues(r Record) []any {
	return Values(r)
}

func (ft ForeignTable) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(ft.prepareInsertStatementString())
}

func (ft ForeignTable) Copy(rs Records) ForeignTable {
	return ForeignTable{Records: rs}
}

func (t DatabaseTable) Name() string {
	rs := t.Records

	if len(rs) > 0 {
		return strings.Split(fmt.Sprintf("%T", rs[0]), ".")[1]
	}

	return unknownTableName
}

func (t DatabaseTable) Fields() []Field {
	rs := t.Records

	if len(rs) > 0 {
		return Fields(rs[0])
	}

	return unknownTableFields
}

func (pt PrimaryTable) prepareInsertStatementString() string {
	tableName := pt.Name()
	tableFields := pt.InsertFields()

	return prepareInsertStatementString(tableName, tableFields)
}

func (pt PrimaryTable) prepareFindStatementString() string {
	tableName := pt.Name()
	tableFindFields := pt.FindFields()

	return prepareFindStatementString(tableName, tableFindFields)
}

func (ft ForeignTable) prepareInsertStatementString() string {
	tableName := ft.Name()
	tableFields := ft.Fields()

	return prepareInsertStatementString(tableName, tableFields)
}

func prepareInsertStatementString(tableName string, tableFields []Field) string {
	tablePlaceholders := stringSliceMap(func(v any) string { return prepareStatementPlaceholderSymbol }, tableFields)

	tableFieldsJoin := strings.Join(toStringSlice(tableFields), prepareStatementMultipleFieldsSeparator)
	tablePlaceholdersJoin := strings.Join(toStringSlice(tablePlaceholders), prepareStatementMultipleFieldsSeparator)

	return fmt.Sprintf(prepareInsertStatementSQLString, tableName, tableFieldsJoin, tablePlaceholdersJoin)
}

func prepareFindStatementString(tableName string, tableFields []Field) string {
	tablePlaceholders := stringSliceMap(func(v any) string { return prepareStatementPlaceholderSymbol }, tableFields)

	tableFieldsJoin := strings.Join(toStringSlice(tableFields), prepareStatementMultipleFieldsSeparator)
	tablePlaceholdersJoin := strings.Join(toStringSlice(tablePlaceholders), prepareStatementMultipleFieldsSeparator)

	return fmt.Sprintf(prepareFindStatementSQLString, tableName, tableFieldsJoin, tablePlaceholdersJoin)
}

func toStringSlice[T any](s []T) []string {
	return stringSliceMap(
		func(v any) string {
			return fmt.Sprintf("%v", v)
		}, s,
	)
}

func stringSliceMap[T any](m func(v any) string, s []T) []string {
	ss := make([]string, len(s))

	for i, v := range s {
		ss[i] = m(v)
	}

	return ss
}
