package entity

type Tuple struct {
	Key   string
	Value any
}

func NewTuple(k string, v any) Tuple {
	return Tuple{
		Key:   k,
		Value: v,
	}
}
