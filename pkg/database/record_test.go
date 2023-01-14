package database

import (
	"reflect"
	"testing"

	"github.com/palavrapasse/damn/pkg/entity"
	. "github.com/palavrapasse/damn/pkg/entity/query"
)

func TestValuesReturnsSchemaValuesIfRecordIsBadActor(t *testing.T) {
	r := BadActor{BaId: 1, Identifier: "l33t"}

	expectedValues := []any{r.BaId, r.Identifier}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsCredentials(t *testing.T) {
	r := Credentials{CredId: 1, Password: "my.password"}

	expectedValues := []any{r.CredId, r.Password}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsHashCredentials(t *testing.T) {
	r := HashCredentials{CredId: 1, HSHA256: entity.NewHSHA256("my.password")}

	expectedValues := []any{r.CredId, r.HSHA256}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsHashUser(t *testing.T) {
	r := HashUser{UserId: 1, HSHA256: entity.NewHSHA256("my.email@gmail.com")}

	expectedValues := []any{r.UserId, r.HSHA256}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakBadActor(t *testing.T) {
	r := LeakBadActor{BaId: 1, LeakId: 2}

	expectedValues := []any{r.BaId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakCredentials(t *testing.T) {
	r := LeakCredentials{CredId: 1, LeakId: 2}

	expectedValues := []any{r.CredId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakPlatform(t *testing.T) {
	r := LeakPlatform{PlatId: 1, LeakId: 2}

	expectedValues := []any{r.PlatId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeak(t *testing.T) {
	r := Leak{LeakId: 1, ShareDateSC: DateInSeconds(2), Context: "twitter breach"}

	expectedValues := []any{r.LeakId, r.ShareDateSC, r.Context}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakUser(t *testing.T) {
	r := LeakUser{UserId: 1, LeakId: 2}

	expectedValues := []any{r.UserId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsPlatform(t *testing.T) {
	r := Platform{PlatId: 1, Name: "twitter"}

	expectedValues := []any{r.PlatId, r.Name}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsUserCredentials(t *testing.T) {
	r := UserCredentials{CredId: 1, UserId: 2}

	expectedValues := []any{r.CredId, r.UserId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsUser(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}

	expectedValues := []any{r.UserId, r.Email}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsBadActor(t *testing.T) {
	r := BadActor{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsCredentials(t *testing.T) {
	r := Credentials{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsLeak(t *testing.T) {
	r := Leak{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsPlatform(t *testing.T) {
	r := Platform{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsUser(t *testing.T) {
	r := User{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsSameRecordIfRecordIsNotPrimary(t *testing.T) {
	r := LeakCredentials{}
	k := entity.AutoGenKey(500)
	expectedRecord := r

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have not set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}
