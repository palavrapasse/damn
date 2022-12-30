package database

import (
	"reflect"
	"testing"

	. "github.com/palavrapasse/damn/pkg/entity"
)

func TestTableNameReturnsUnknownTableNameIfItDoesNotContainAnyRecord(t *testing.T) {
	tb := DatabaseTable{}

	name := tb.Name()

	if name != unknownTableName {
		t.Fatalf("Table does not contain any record, so it's name can't be inferred, but got: %s\n", name)
	}
}

func TestBadActorTableNameReturnsBadActor(t *testing.T) {
	tb := NewBadActorTable([]BadActor{{}})
	expectedTableName := "BadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("BadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestCredentialsTableNameReturnsCredentials(t *testing.T) {
	tb := NewCredentialsTable([]Credentials{{}})
	expectedTableName := "Credentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Credentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashCredentialsTableNameReturnsHashCredentials(t *testing.T) {
	tb := NewHashCredentialsTable([]Credentials{{}})
	expectedTableName := "HashCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashUserTableNameReturnsHashUser(t *testing.T) {
	tb := NewHashUserTable([]User{{}})
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
	tb := NewPlatformTable([]Platform{{}})
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
	tb := NewUserTable([]User{{}})
	expectedTableName := "User"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestTableFieldsReturnsUnknownTableIfItDoesNotContainAnyRecord(t *testing.T) {
	tb := DatabaseTable{}

	fields := tb.Fields()

	if !reflect.DeepEqual(fields, unknownTableFields) {
		t.Fatalf("Table does not contain any record, so it's fields can't be inferred, but got: %v\n", fields)
	}
}

func TestPrimaryTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records{User{UserId: 1}}
	newRecords := Records{User{UserId: 2}}

	oldTable := UserTable{Records: oldRecords}
	expectedNewTable := UserTable{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestForeignTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records{LeakUser{UserId: 1}}
	newRecords := Records{LeakUser{UserId: 2}}

	oldTable := LeakUserTable{Records: oldRecords}
	expectedNewTable := LeakUserTable{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestToBadActorSliceReturnsTableBadActorRecords(t *testing.T) {
	expectedBadActorSlice := []BadActor{{BaId: 1}, {BaId: 2}}

	bat := NewBadActorTable(expectedBadActorSlice)
	bas := bat.ToBadActorSlice()

	if !reflect.DeepEqual(bas, expectedBadActorSlice) {
		t.Fatalf("ToBadActorSlice should have return a slice with all BadActor records, but got: %v", bas)
	}
}

func TestToCredentialsSliceReturnsTableCredentialsRecords(t *testing.T) {
	expectedCredentialsSlice := []Credentials{{CredId: 1}, {CredId: 2}}

	bat := NewCredentialsTable(expectedCredentialsSlice)
	bas := bat.ToCredentialsSlice()

	if !reflect.DeepEqual(bas, expectedCredentialsSlice) {
		t.Fatalf("ToCredentialsSlice should have return a slice with all Credentials records, but got: %v", bas)
	}
}

func TestToLeakSliceReturnsTableLeakRecords(t *testing.T) {
	expectedLeakSlice := []Leak{{LeakId: 1}, {LeakId: 2}}

	bat := NewLeakTable(expectedLeakSlice...)
	bas := bat.ToLeakSlice()

	if !reflect.DeepEqual(bas, expectedLeakSlice) {
		t.Fatalf("ToLeakSlice should have return a slice with all Leak records, but got: %v", bas)
	}
}

func TestToPlatformSliceReturnsTablePlatformRecords(t *testing.T) {
	expectedPlatformSlice := []Platform{{PlatId: 1}, {PlatId: 2}}

	bat := NewPlatformTable(expectedPlatformSlice)
	bas := bat.ToPlatformSlice()

	if !reflect.DeepEqual(bas, expectedPlatformSlice) {
		t.Fatalf("ToPlatformSlice should have return a slice with all Platform records, but got: %v", bas)
	}
}

func TestToUserSliceReturnsTableUserRecords(t *testing.T) {
	expectedUserSlice := []User{{UserId: 1}, {UserId: 2}}

	bat := NewUserTable(expectedUserSlice)
	bas := bat.ToUserSlice()

	if !reflect.DeepEqual(bas, expectedUserSlice) {
		t.Fatalf("ToUserSlice should have return a slice with all User records, but got: %v", bas)
	}
}

func TestInsertFieldsReturnsPrimaryTableFieldsExceptPrimaryKey(t *testing.T) {
	ut := NewUserTable([]User{{}})
	expectedFields := []Field{"email"}

	insertFields := ut.InsertFields()

	if !reflect.DeepEqual(insertFields, expectedFields) {
		t.Fatalf("InsertFields should have return a slice with table fields except primary key, but got: %v", insertFields)
	}
}

func TestInsertFieldsReturnsAllForeignTableFields(t *testing.T) {
	ut := NewHashUserTable([]User{{}})
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
	ut := NewHashUserTable([]User{r})
	expectedValues := []any{r.UserId, r.Email}

	insertValues := ut.InsertValues(r)

	if !reflect.DeepEqual(insertValues, expectedValues) {
		t.Fatalf("InsertValues should have return a slice with all record values, but got: %v", insertValues)
	}
}

func TestTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	bat := NewBadActorTable([]BadActor{{}}).prepareInsertStatementString()
	crt := NewCredentialsTable([]Credentials{{}}).prepareInsertStatementString()
	hct := NewHashCredentialsTable([]Credentials{{}}).prepareInsertStatementString()
	hut := NewHashUserTable([]User{{}}).prepareInsertStatementString()
	lbat := NewLeakBadActorTable(map[Leak][]BadActor{{}: {BadActor{}}}).prepareInsertStatementString()
	lct := NewLeakCredentialsTable(map[Leak][]Credentials{{}: {Credentials{}}}).prepareInsertStatementString()
	lpt := NewLeakPlatformTable(map[Leak][]Platform{{}: {Platform{}}}).prepareInsertStatementString()
	lt := NewLeakTable(Leak{}).prepareInsertStatementString()
	lut := NewLeakUserTable(map[Leak][]User{{}: {User{}}}).prepareInsertStatementString()
	pt := NewPlatformTable([]Platform{{}}).prepareInsertStatementString()
	uct := NewUserCredentialsTable(map[User]Credentials{{}: {}}).prepareInsertStatementString()
	ut := NewUserTable([]User{{}}).prepareInsertStatementString()

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

func TestTablePrepareInsertStatementReturnsSchemaFindStatement(t *testing.T) {
	bat := NewBadActorTable([]BadActor{{}}).prepareFindStatementString()
	crt := NewCredentialsTable([]Credentials{{}}).prepareFindStatementString()
	lt := NewLeakTable(Leak{}).prepareFindStatementString()
	pt := NewPlatformTable([]Platform{{}}).prepareFindStatementString()
	ut := NewUserTable([]User{{}}).prepareFindStatementString()

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
