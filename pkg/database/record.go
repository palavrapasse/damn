package database

import (
	"reflect"

	"github.com/palavrapasse/damn/pkg/entity"
)

type Record interface{}
type Records[R Record] []R

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

func CopyWithNewKey[R Record](r R, k entity.AutoGenKey) R {
	rr := r

	copy := reflect.ValueOf(r).MethodByName("Copy")

	if copy.IsValid() {
		rr = copy.Call([]reflect.Value{reflect.ValueOf(k)})[0].Interface().(R)
	}

	return rr
}

func reflectRecord(r Record) []entity.Tuple {
	return reflect.ValueOf(r).MethodByName("Record").Call(nil)[0].Interface().([]entity.Tuple)
}
