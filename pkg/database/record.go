package database

import (
	"reflect"

	"github.com/palavrapasse/damn/pkg/entity"
)

type Record interface{}
type Records []Record

type Field string

func Fields(r Record) []Field {
	rr := reflectRecord(r)
	fs := make([]Field, len(rr))

	for i, f := range rr {
		fs[i] = Field(f.Key)
	}

	return fs
}

func Values(r Record) []any {
	rr := reflectRecord(r)
	vs := make([]any, len(rr))

	for i, f := range rr {
		vs[i] = f.Value
	}

	return vs
}

func CopyWithNewKey(r Record, k entity.AutoGenKey) Record {
	var rr Record

	switch r.(type) {
	case entity.BadActor:
		r := r.(entity.BadActor)
		rr = r.Copy(k)
	case entity.Credentials:
		r := r.(entity.Credentials)
		rr = r.Copy(k)
	case entity.Leak:
		r := r.(entity.Leak)
		rr = r.Copy(k)
	case entity.Platform:
		r := r.(entity.Platform)
		rr = r.Copy(k)
	case entity.User:
		r := r.(entity.User)
		rr = r.Copy(k)
	default:
		rr = r
	}

	return rr
}

func reflectRecord(r Record) []entity.Tuple {
	return reflect.ValueOf(r).MethodByName("Record").Call(nil)[0].Interface().([]entity.Tuple)
}
