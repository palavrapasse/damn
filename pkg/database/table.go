package database

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
	"sync"

	. "github.com/palavrapasse/damn/pkg/entity"
	. "github.com/palavrapasse/damn/pkg/entity/query"
	. "github.com/palavrapasse/damn/pkg/entity/subscribe"
)

const (
	prepareInsertStatementSQLString         = "INSERT OR IGNORE INTO %s (%s) VALUES (%s)"
	prepareFindStatementSQLString           = "SELECT * FROM %s WHERE (%s) = (%s) LIMIT 1"
	prepareStatementPlaceholderSymbol       = "?"
	prepareStatementMultipleFieldsSeparator = ", "
)

const (
	AscendingSortOrderKeyword  = "ASC"
	DescendingSortOrderKeyword = "DESC"
)

type Table[R Record] interface {
	Name() string
	Records() []R
	Fields() []Field
	Values(R) []any
	Copy([]R) DatabaseTable[R]
	Generalize() DatabaseTable[Record]
	HasPrimaryKeySet(R) bool
	InsertFields() []Field
	InsertValues(R) []any
	FindFields() []Field
	FindValues(R) []any
	PrepareInsertStatement(*sql.Tx) (*sql.Stmt, error)
	PrepareFindStatement(*sql.Tx) (*sql.Stmt, error)
}

type DatabaseTable[R Record] struct {
	Table[R]
	Records[R]
}

type PrimaryTable[R Record] DatabaseTable[R]
type ForeignTable[R Record] DatabaseTable[R]

type concurrentHashForeignTableResult[R Record] struct {
	hashForeignTable ForeignTable[R]
	routineId        int
}

type concurrentPrimaryTableResult[R Record] struct {
	primaryTable PrimaryTable[R]
	routineId    int
}

func MultiplePlaceholder(lv int) string {
	phs := make([]string, lv)

	for i := 0; i < lv; i++ {
		phs[i] = prepareStatementPlaceholderSymbol
	}

	return strings.Join(phs, prepareStatementMultipleFieldsSeparator)
}

func NewBadActorTable(ba []BadActor) PrimaryTable[BadActor] {
	return PrimaryTable[BadActor]{
		Records: ba,
	}
}

func NewCredentialsTable(cr []Credentials) PrimaryTable[Credentials] {
	return PrimaryTable[Credentials]{
		Records: cr,
	}
}

func NewLeakTable(ls ...Leak) PrimaryTable[Leak] {
	return PrimaryTable[Leak]{
		Records: ls,
	}
}

func NewPlatformTable(ps []Platform) PrimaryTable[Platform] {
	return PrimaryTable[Platform]{
		Records: ps,
	}
}

func NewUserTable(us []User) PrimaryTable[User] {
	return PrimaryTable[User]{
		Records: us,
	}
}

func NewSubscriberTable(su ...Subscriber) PrimaryTable[Subscriber] {
	return PrimaryTable[Subscriber]{
		Records: su,
	}
}

func NewAffectedTable(af []Affected) PrimaryTable[Affected] {
	return PrimaryTable[Affected]{
		Records: af,
	}
}

func NewHashCredentialsTable(cr []Credentials) ForeignTable[HashCredentials] {
	rs := make(Records[HashCredentials], len(cr))

	for i, v := range cr {
		rs[i] = NewHashCredentials(v)
	}

	return ForeignTable[HashCredentials]{
		Records: rs,
	}
}

func NewHashUserTable(us []User) ForeignTable[HashUser] {
	rs := make(Records[HashUser], len(us))

	for i, v := range us {
		rs[i] = NewHashUser(v)
	}

	return ForeignTable[HashUser]{
		Records: rs,
	}
}

func NewLeakBadActorTable(lba map[Leak][]BadActor) ForeignTable[LeakBadActor] {
	rs := Records[LeakBadActor]{}

	for l, bas := range lba {
		for _, ba := range bas {
			rs = append(rs, NewLeakBadActor(ba, l))
		}
	}

	return ForeignTable[LeakBadActor]{
		Records: rs,
	}
}

func NewLeakCredentialsTable(lcr map[Leak][]Credentials) ForeignTable[LeakCredentials] {
	rs := Records[LeakCredentials]{}

	for l, crs := range lcr {
		for _, cr := range crs {
			rs = append(rs, NewLeakCredentials(cr, l))
		}
	}

	return ForeignTable[LeakCredentials]{
		Records: rs,
	}
}

func NewLeakPlatformTable(lpt map[Leak][]Platform) ForeignTable[LeakPlatform] {
	rs := Records[LeakPlatform]{}

	for l, pts := range lpt {
		for _, pt := range pts {
			rs = append(rs, NewLeakPlatform(pt, l))
		}
	}

	return ForeignTable[LeakPlatform]{
		Records: rs,
	}
}

func NewLeakUserTable(lus map[Leak][]User) ForeignTable[LeakUser] {
	rs := Records[LeakUser]{}

	for l, us := range lus {
		for _, u := range us {
			rs = append(rs, NewLeakUser(u, l))
		}
	}

	return ForeignTable[LeakUser]{
		Records: rs,
	}
}

func NewUserCredentialsTable(uc map[User]Credentials) ForeignTable[UserCredentials] {
	rs := make(Records[UserCredentials], len(uc))

	i := 0

	for u, c := range uc {
		rs[i] = UserCredentials{CredId: c.CredId, UserId: u.UserId}

		i++
	}

	return ForeignTable[UserCredentials]{
		Records: rs,
	}
}

func NewSubscriberAffectedTable(s Subscriber, a []Affected) ForeignTable[SubscriberAffected] {
	rs := make(Records[SubscriberAffected], len(a))

	for i, aff := range a {
		rs[i] = SubscriberAffected{AffId: aff.AffectedId, SubId: s.SubscriberId}
	}

	return ForeignTable[SubscriberAffected]{
		Records: rs,
	}
}

func NewConcurrentHashForeignTable[F HashCredentials | HashUser, P Credentials | User](maxElementsOfGoroutine int, primaryElements []P, newForeignTableCallback func([]P) ForeignTable[F]) ForeignTable[F] {

	ngoroutines := 1
	nelements := len(primaryElements)

	if nelements > maxElementsOfGoroutine {
		ngoroutines = int(math.Ceil(float64(nelements) / float64(maxElementsOfGoroutine)))
	}

	resultChan := make(chan concurrentHashForeignTableResult[F])

	var wg sync.WaitGroup

	wg.Add(ngoroutines)

	for i := 0; i < ngoroutines; i++ {

		init := i * maxElementsOfGoroutine
		end := (i + 1) * maxElementsOfGoroutine
		if end > nelements {
			end = nelements
		}

		go func(lines []P, routineId int) {

			defer wg.Done()
			resultChan <- concurrentHashForeignTableResult[F]{
				routineId:        routineId,
				hashForeignTable: newForeignTableCallback(lines),
			}

		}(primaryElements[init:end], i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	mapResult := make(map[int]concurrentHashForeignTableResult[F])

	for r := range resultChan {
		mapResult[r.routineId] = r
	}

	result := ForeignTable[F]{}

	for i := 0; i < ngoroutines; i++ {
		result.Records = append(result.Records, mapResult[i].hashForeignTable.Records...)
	}

	return result
}

func NewConcurrentPrimaryTable[P BadActor | User | Credentials | Platform](maxElementsOfGoroutine int, primaryElements []P, newPrimaryTableCallback func([]P) PrimaryTable[P]) PrimaryTable[P] {

	ngoroutines := 1
	nelements := len(primaryElements)

	if nelements > maxElementsOfGoroutine {
		ngoroutines = int(math.Ceil(float64(nelements) / float64(maxElementsOfGoroutine)))
	}

	resultChan := make(chan concurrentPrimaryTableResult[P])

	var wg sync.WaitGroup

	wg.Add(ngoroutines)

	for i := 0; i < ngoroutines; i++ {

		init := i * maxElementsOfGoroutine
		end := (i + 1) * maxElementsOfGoroutine
		if end > nelements {
			end = nelements
		}

		go func(lines []P, routineId int) {

			defer wg.Done()
			resultChan <- concurrentPrimaryTableResult[P]{
				routineId:    routineId,
				primaryTable: newPrimaryTableCallback(lines),
			}

		}(primaryElements[init:end], i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	mapResult := make(map[int]concurrentPrimaryTableResult[P])

	for r := range resultChan {
		mapResult[r.routineId] = r
	}

	result := PrimaryTable[P]{}

	for i := 0; i < ngoroutines; i++ {
		result.Records = append(result.Records, mapResult[i].primaryTable.Records...)
	}

	return result
}

func (pt PrimaryTable[R]) Name() string {
	return DatabaseTable[R](pt).Name()
}

func (pt PrimaryTable[R]) Fields() []Field {
	return DatabaseTable[R](pt).Fields()
}

func (pt PrimaryTable[R]) Values(r R) []any {
	return Values(r)
}

func (pt PrimaryTable[R]) InsertFields() []Field {
	return DatabaseTable[R](pt).Fields()[1:]
}

func (pt PrimaryTable[R]) InsertValues(r R) []any {
	return Values(r)[1:]
}

func (pt PrimaryTable[R]) FindFields() []Field {
	// todo: rely on sql tags
	return DatabaseTable[R](pt).Fields()[1:]
}

func (pt PrimaryTable[R]) FindValues(r R) []any {
	// todo: rely on sql tags
	return Values(r)[1:]
}

func (pt PrimaryTable[R]) HasPrimaryKeySet(r R) bool {
	// todo: rely on sql tags
	return Values(r)[0] != AutoGenKey(0)
}

func (pt PrimaryTable[R]) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(pt.prepareInsertStatementString())

}

func (pt PrimaryTable[R]) PrepareFindStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(pt.prepareFindStatementString())

}

func (pt PrimaryTable[R]) Copy(rs Records[R]) PrimaryTable[R] {
	return PrimaryTable[R]{Records: rs}
}

func (ft ForeignTable[R]) Name() string {
	return DatabaseTable[R](ft).Name()
}

func (ft ForeignTable[R]) Fields() []Field {
	return DatabaseTable[R](ft).Fields()
}

func (ft ForeignTable[R]) InsertFields() []Field {
	return DatabaseTable[R](ft).Fields()
}

func (ft ForeignTable[R]) InsertValues(r R) []any {
	return Values(r)
}

func (ft ForeignTable[R]) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(ft.prepareInsertStatementString())
}

func (ft ForeignTable[R]) Copy(rs Records[R]) ForeignTable[R] {
	return ForeignTable[R]{Records: rs}
}

func (t DatabaseTable[R]) Name() string {
	return strings.Split(fmt.Sprintf("%T", new(R)), ".")[1]
}

func (t DatabaseTable[R]) Fields() []Field {
	rs := t.Records
	var fs []Field

	if len(rs) > 0 {
		fs = Fields(rs[0])
	} else {
		fs = Fields(new(R))
	}

	return fs
}

func (pt PrimaryTable[R]) prepareInsertStatementString() string {
	tableName := pt.Name()
	tableFields := pt.InsertFields()

	return prepareInsertStatementString(tableName, tableFields)
}

func (pt PrimaryTable[R]) prepareFindStatementString() string {
	tableName := pt.Name()
	tableFindFields := pt.FindFields()

	return prepareFindStatementString(tableName, tableFindFields)
}

func (ft ForeignTable[R]) prepareInsertStatementString() string {
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
