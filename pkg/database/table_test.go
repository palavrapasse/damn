package database

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/palavrapasse/damn/pkg/entity/query"
	. "github.com/palavrapasse/damn/pkg/entity/subscribe"
)

func TestMultiplePlaceholderReturnsEmptyStringIfValueCountIsZero(t *testing.T) {
	lv := 0

	mph := MultiplePlaceholder(lv)
	emph := ""

	if mph != emph {
		t.Fatalf("function should have returned (%s), but got: (%s)", mph, emph)
	}
}

func TestMultiplePlaceholderReturnsSinglePlaceholderIfValueCountIsOne(t *testing.T) {
	lv := 1

	mph := MultiplePlaceholder(lv)
	emph := prepareStatementPlaceholderSymbol

	if mph != emph {
		t.Fatalf("function should have returned (%s), but got: (%s)", mph, emph)
	}
}

func TestMultiplePlaceholderReturnsMultiplePlaceholderIfValueCountIsMoreThanOne(t *testing.T) {
	lv := 3

	mph := MultiplePlaceholder(lv)
	emph := fmt.Sprintf("%s, %s, %s", prepareStatementPlaceholderSymbol, prepareStatementPlaceholderSymbol, prepareStatementPlaceholderSymbol)

	if mph != emph {
		t.Fatalf("function should have returned (%s), but got: (%s)", mph, emph)
	}
}

func TestBadActorTableNameReturnsBadActor(t *testing.T) {
	tb := NewBadActorTable([]BadActor{})
	expectedTableName := "BadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("BadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestCredentialsTableNameReturnsCredentials(t *testing.T) {
	tb := NewCredentialsTable([]Credentials{})
	expectedTableName := "Credentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Credentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashCredentialsTableNameReturnsHashCredentials(t *testing.T) {
	tb := NewHashCredentialsTable([]Credentials{})
	expectedTableName := "HashCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashUserTableNameReturnsHashUser(t *testing.T) {
	tb := NewHashUserTable([]User{})
	expectedTableName := "HashUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakBadActorTableNameReturnsLeakBadActor(t *testing.T) {
	tb := NewLeakBadActorTable(map[Leak][]BadActor{{}: {BadActor{}}})
	expectedTableName := "LeakBadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakBadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakCredentialsTableNameReturnsLeakCredentials(t *testing.T) {
	tb := NewLeakCredentialsTable(map[Leak][]Credentials{{}: {Credentials{}}})
	expectedTableName := "LeakCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakPlatformTableNameReturnsLeakPlatform(t *testing.T) {
	tb := NewLeakPlatformTable(map[Leak][]Platform{{}: {Platform{}}})
	expectedTableName := "LeakPlatform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakPlatform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakTableNameReturnsLeak(t *testing.T) {
	tb := NewLeakTable(Leak{})
	expectedTableName := "Leak"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Leak table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakUserTableNameReturnsLeakUser(t *testing.T) {
	tb := NewLeakUserTable(map[Leak][]User{{}: {User{}}})
	expectedTableName := "LeakUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestPlatformTableNameReturnsPlatform(t *testing.T) {
	tb := NewPlatformTable([]Platform{})
	expectedTableName := "Platform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Platform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserCredentialsTableNameReturnsUserCredentials(t *testing.T) {
	tb := NewUserCredentialsTable(map[User]Credentials{{}: {}})
	expectedTableName := "UserCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("UserCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserTableNameReturnsUser(t *testing.T) {
	tb := NewUserTable([]User{})
	expectedTableName := "User"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestSubscriberTableNameReturnsSubscriber(t *testing.T) {
	tb := NewSubscriberTable(Subscriber{})
	expectedTableName := "Subscriber"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestAffectedTableNameReturnsAffected(t *testing.T) {
	tb := NewAffectedTable([]Affected{})
	expectedTableName := "Affected"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestSubscriberAffectedTableNameReturnsSubscriberAffected(t *testing.T) {
	tb := NewSubscriberAffectedTable(Subscriber{}, []Affected{})
	expectedTableName := "SubscriberAffected"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestPrimaryTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records[User]{User{UserId: 1}}
	newRecords := Records[User]{User{UserId: 2}}

	oldTable := PrimaryTable[User]{Records: oldRecords}
	expectedNewTable := PrimaryTable[User]{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestForeignTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records[LeakUser]{LeakUser{UserId: 1}}
	newRecords := Records[LeakUser]{LeakUser{UserId: 2}}

	oldTable := ForeignTable[LeakUser]{Records: oldRecords}
	expectedNewTable := ForeignTable[LeakUser]{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestInsertFieldsReturnsPrimaryTableFieldsExceptPrimaryKey(t *testing.T) {
	ut := NewUserTable([]User{})
	expectedFields := []Field{"email"}

	insertFields := ut.InsertFields()

	if !reflect.DeepEqual(insertFields, expectedFields) {
		t.Fatalf("InsertFields should have return a slice with table fields except primary key, but got: %v", insertFields)
	}
}

func TestInsertFieldsReturnsAllForeignTableFields(t *testing.T) {
	ut := NewHashUserTable([]User{})
	expectedFields := []Field{"userid", "hsha256"}

	insertFields := ut.InsertFields()

	if !reflect.DeepEqual(insertFields, expectedFields) {
		t.Fatalf("InsertFields should have return a slice with all table fields, but got: %v", insertFields)
	}
}

func TestInsertValuesReturnsPrimaryTableValuesExceptPrimaryKeyValue(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}
	ut := NewUserTable([]User{r})
	expectedValues := []any{r.Email}

	insertValues := ut.InsertValues(r)

	if !reflect.DeepEqual(insertValues, expectedValues) {
		t.Fatalf("InsertValues should have return a slice with record values except primary key, but got: %v", insertValues)
	}
}

func TestInsertValuesReturnsAllForeignTableValues(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}
	hu := NewHashUser(r)
	hut := ForeignTable[HashUser]{Records: []HashUser{hu}}
	expectedValues := []any{hu.UserId, hu.HSHA256}

	insertValues := hut.InsertValues(NewHashUser(r))

	if !reflect.DeepEqual(insertValues, expectedValues) {
		t.Fatalf("InsertValues should have return a slice with all record values, but got: %v", insertValues)
	}
}

func TestTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	bat := PrimaryTable[BadActor]{}.prepareInsertStatementString()
	crt := PrimaryTable[Credentials]{}.prepareInsertStatementString()
	hct := ForeignTable[HashCredentials]{}.prepareInsertStatementString()
	hut := ForeignTable[HashUser]{}.prepareInsertStatementString()
	lbat := ForeignTable[LeakBadActor]{}.prepareInsertStatementString()
	lct := ForeignTable[LeakCredentials]{}.prepareInsertStatementString()
	lpt := ForeignTable[LeakPlatform]{}.prepareInsertStatementString()
	lt := PrimaryTable[Leak]{}.prepareInsertStatementString()
	lut := ForeignTable[LeakUser]{}.prepareInsertStatementString()
	pt := PrimaryTable[Platform]{}.prepareInsertStatementString()
	uct := ForeignTable[UserCredentials]{}.prepareInsertStatementString()
	ut := PrimaryTable[User]{}.prepareInsertStatementString()

	tableInsertSchemaMapping := map[string]string{
		bat:  "INSERT OR IGNORE INTO BadActor (identifier) VALUES (?)",
		crt:  "INSERT OR IGNORE INTO Credentials (password) VALUES (?)",
		hct:  "INSERT OR IGNORE INTO HashCredentials (credid, hsha256) VALUES (?, ?)",
		hut:  "INSERT OR IGNORE INTO HashUser (userid, hsha256) VALUES (?, ?)",
		lbat: "INSERT OR IGNORE INTO LeakBadActor (baid, leakid) VALUES (?, ?)",
		lct:  "INSERT OR IGNORE INTO LeakCredentials (credid, leakid) VALUES (?, ?)",
		lpt:  "INSERT OR IGNORE INTO LeakPlatform (platid, leakid) VALUES (?, ?)",
		lt:   "INSERT OR IGNORE INTO Leak (sharedatesc, context) VALUES (?, ?)",
		lut:  "INSERT OR IGNORE INTO LeakUser (userid, leakid) VALUES (?, ?)",
		pt:   "INSERT OR IGNORE INTO Platform (name) VALUES (?)",
		uct:  "INSERT OR IGNORE INTO UserCredentials (credid, userid) VALUES (?, ?)",
		ut:   "INSERT OR IGNORE INTO User (email) VALUES (?)",
	}

	for ts, s := range tableInsertSchemaMapping {
		expectedSchema := s
		statement := ts

		if statement != expectedSchema {
			t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", statement)
		}
	}
}

func TestTablePrepareFindStatementReturnsSchemaFindStatement(t *testing.T) {
	bat := PrimaryTable[BadActor]{}.prepareFindStatementString()
	crt := PrimaryTable[Credentials]{}.prepareFindStatementString()
	lt := PrimaryTable[Leak]{}.prepareFindStatementString()
	pt := PrimaryTable[Platform]{}.prepareFindStatementString()
	ut := PrimaryTable[User]{}.prepareFindStatementString()

	tableFindSchemaMapping := map[string]string{
		bat: "SELECT * FROM BadActor WHERE (identifier) = (?) LIMIT 1",
		crt: "SELECT * FROM Credentials WHERE (password) = (?) LIMIT 1",
		lt:  "SELECT * FROM Leak WHERE (sharedatesc, context) = (?, ?) LIMIT 1",
		pt:  "SELECT * FROM Platform WHERE (name) = (?) LIMIT 1",
		ut:  "SELECT * FROM User WHERE (email) = (?) LIMIT 1",
	}

	for ts, s := range tableFindSchemaMapping {
		expectedSchema := s
		statement := ts

		if statement != expectedSchema {
			t.Fatalf("Prepared find statement should be the same as defined in the schema, but got: %v", statement)
		}
	}
}

func TestPrimaryTableHasPrimaryKeySetReturnsTrueIfAutoGenKeyValueIsDifferentThanZero(t *testing.T) {
	r := User{UserId: 1}
	tb := NewUserTable([]User{r})

	expectedVerification := true

	hasPrimaryKeySet := tb.HasPrimaryKeySet(r)

	if hasPrimaryKeySet != expectedVerification {
		t.Fatalf("Record has its primary key with a value greater than 0, but HasPrimaryKey returned: %v", hasPrimaryKeySet)
	}
}

func TestPrimaryTableHasPrimaryKeySetReturnsFalseIfAutoGenKeyValueIsEqualToZero(t *testing.T) {
	r := User{UserId: 1}
	tb := NewUserTable([]User{r})

	expectedVerification := true

	hasPrimaryKeySet := tb.HasPrimaryKeySet(r)

	if hasPrimaryKeySet != expectedVerification {
		t.Fatalf("Record has its primary key with a value greater than 0, but HasPrimaryKey returned: %v", hasPrimaryKeySet)
	}
}

func TestPrimaryTableValuesReturnsAllRecordValues(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}
	tb := NewUserTable([]User{r})

	expectedValues := []any{r.UserId, r.Email}

	values := tb.Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have returned all values, but got: %v", values...)
	}
}

func TestPrimaryTableFindValuesReturnsAllRecordValuesExceptPrimaryKey(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}
	tb := NewUserTable([]User{r})

	expectedValues := []any{r.Email}

	findValues := tb.FindValues(r)

	if !reflect.DeepEqual(findValues, expectedValues) {
		t.Fatalf("FindValues should have returned all values except primary key, but got: %v", findValues...)
	}
}

func TestPrimaryTableFieldsReturnsAllTableFieldsNames(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}
	tb := NewUserTable([]User{r})

	expectedFields := []Field{"userid", "email"}

	fields := tb.Fields()

	if !reflect.DeepEqual(fields, expectedFields) {
		t.Fatalf("Fields should have returned all field names, but got: %v", fields)
	}
}

func TestNewConcurrentHashForeignTableWithoutGoRoutinesReturnsTheSameAsNoRoutineFunction(t *testing.T) {
	tb := []Credentials{{Password: Password("pass1word")}, {Password: Password("pass2word")}, {Password: Password("pass3word")}, {Password: Password("pass3word")}}

	expected := NewHashCredentialsTable(tb)

	result := NewConcurrentHashForeignTable(len(tb), tb, NewHashCredentialsTable)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("NewConcurrentHashForeignTable should have returned the same result and order as the function with no goroutines, but got: %v", result)
	}
}

func TestNewConcurrentHashForeignTableWithGoRoutinesReturnsTheSameAsNoRoutineFunction(t *testing.T) {
	tb := []Credentials{{Password: Password("pass1word")}, {Password: Password("pass2word")}, {Password: Password("pass3word")}, {Password: Password("pass4word")}, {Password: Password("pass5word")}}

	expected := NewHashCredentialsTable(tb)

	result := NewConcurrentHashForeignTable(len(tb)-2, tb, NewHashCredentialsTable)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("NewConcurrentHashForeignTable should have returned the same result and order as the function with no goroutines, but got: %v", result)
	}
}

func TestNewConcurrentPrimaryTableWithoutGoRoutinesReturnsTheSameAsNoRoutineFunction(t *testing.T) {
	tb := []Credentials{{Password: Password("pass1word")}, {Password: Password("pass2word")}, {Password: Password("pass3word")}, {Password: Password("pass3word")}}

	expected := NewHashCredentialsTable(tb)

	result := NewConcurrentPrimaryTable(len(tb), tb, NewCredentialsTable)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("NewConcurrentPrimaryTable should have returned the same result and order as the function with no goroutines, but got: %v", result)
	}
}

func TestNewConcurrentPrimaryTableWithGoRoutinesReturnsTheSameAsNoRoutineFunction(t *testing.T) {
	tb := []Credentials{{Password: Password("pass1word")}, {Password: Password("pass2word")}, {Password: Password("pass3word")}, {Password: Password("pass4word")}, {Password: Password("pass5word")}}

	expected := NewHashCredentialsTable(tb)

	result := NewConcurrentPrimaryTable(len(tb)-2, tb, NewCredentialsTable)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("NewConcurrentPrimaryTable should have returned the same result and order as the function with no goroutines, but got: %v", result)
	}
}
