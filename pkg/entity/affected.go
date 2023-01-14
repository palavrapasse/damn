package entity

const (
	AffectedIdField        = "affid"
	AffectedHashEmailField = "hsha256email"
)

type Affected struct {
	HSHA256Email HSHA256
	AffectedId   AutoGenKey
}

func NewAffected(email string) Affected {
	return Affected{
		HSHA256Email: NewHSHA256(email),
	}
}

func (a Affected) Copy(key AutoGenKey) Affected {
	return Affected{
		AffectedId:   key,
		HSHA256Email: a.HSHA256Email,
	}
}

func (a Affected) Record() []Tuple {
	return []Tuple{
		NewTuple(AffectedIdField, a.AffectedId),
		NewTuple(AffectedHashEmailField, a.HSHA256Email),
	}
}
