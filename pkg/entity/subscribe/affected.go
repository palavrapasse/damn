package subscribe

import "github.com/palavrapasse/damn/pkg/entity"

const (
	AffectedIdField        = "affid"
	AffectedHashEmailField = "hsha256email"
)

type Affected struct {
	HSHA256Email entity.HSHA256
	AffectedId   entity.AutoGenKey
}

func NewAffected(email string) Affected {
	return Affected{
		HSHA256Email: entity.NewHSHA256(email),
	}
}

func (a Affected) Copy(key entity.AutoGenKey) Affected {
	return Affected{
		AffectedId:   key,
		HSHA256Email: a.HSHA256Email,
	}
}

func (a Affected) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(AffectedIdField, a.AffectedId),
		entity.NewTuple(AffectedHashEmailField, a.HSHA256Email),
	}
}
