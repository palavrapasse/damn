package entity

type AutoGenKey int64

func NewAutoGenKey(key int64) AutoGenKey {
	return AutoGenKey(key)
}
