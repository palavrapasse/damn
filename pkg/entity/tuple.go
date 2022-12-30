package entity

type Tuple struct {
	Value any
	Key   string
}

func NewTuple(k string, v any) Tuple {
	return Tuple{
		Key:   k,
		Value: v,
	}
}
