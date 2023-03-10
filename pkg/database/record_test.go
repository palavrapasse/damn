package database

import (
	"reflect"
	"testing"

	"github.com/palavrapasse/damn/pkg/entity"
	. "github.com/palavrapasse/damn/pkg/entity/query"
	. "github.com/palavrapasse/damn/pkg/entity/subscribe"
)

func TestValuesReturnsSchemaValuesIfRecordIsBadActor(t *testing.T) {
	r := BadActor{BaId: 1, Identifier: "l33t"}

	expectedValues := []any{r.BaId, r.Identifier}

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

func TestValuesReturnsSchemaValuesIfRecordIsUser(t *testing.T) {
	r := User{UserId: 1, Email: "my.email@gmail.com"}

	expectedValues := []any{r.UserId, r.Email}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsSubscriber(t *testing.T) {
	r := Subscriber{SubscriberId: 1, B64Email: "base64Email"}

	expectedValues := []any{r.SubscriberId, r.B64Email}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsAffected(t *testing.T) {
	r := Affected{AffectedId: 1, HSHA256Email: entity.NewHSHA256("Email")}

	expectedValues := []any{r.AffectedId, r.HSHA256Email}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsSubscriberAffected(t *testing.T) {
	r := SubscriberAffected{AffId: 1, SubId: 2}

	expectedValues := []any{r.AffId, r.SubId}

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

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsSubscriber(t *testing.T) {
	r := Subscriber{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsAffected(t *testing.T) {
	r := Affected{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}
