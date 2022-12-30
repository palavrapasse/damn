package database

import (
	"reflect"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

type Record interface{}
type Records []Record

type Field string

func Fields(r Record) []Field {
	rf := reflectFields(r)
	fs := make([]Field, len(rf))

	for i, f := range rf {
		fs[i] = Field(strings.ToLower(f.Name))
	}

	return fs
}

func Values(r Record) []any {
	rf := reflectFields(r)
	rv := reflect.ValueOf(r)
	vs := make([]any, len(rf))

	for i, f := range rf {
		vs[i] = rv.FieldByName(f.Name).Interface()
	}

	return vs
}

func CopyWithNewKey(r Record, k entity.AutoGenKey) Record {
	var rr Record

	switch r.(type) {
	case entity.BadActor:
		r, _ := r.(entity.BadActor)
		rr = r.Copy(k)
	case entity.Credentials:
		r, _ := r.(entity.Credentials)
		rr = r.Copy(k)
	case entity.Leak:
		r, _ := r.(entity.Leak)
		rr = r.Copy(k)
	case entity.Platform:
		r, _ := r.(entity.Platform)
		rr = r.Copy(k)
	case entity.User:
		r, _ := r.(entity.User)
		rr = r.Copy(k)
	default:
		rr = r
	}

	return rr
}

func reflectFields(r Record) []reflect.StructField {
	return reflect.VisibleFields(reflect.TypeOf(r))
}
