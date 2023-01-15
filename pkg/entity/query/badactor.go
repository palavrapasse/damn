package query

import (
	"errors"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	BadActorIdField         = "baid"
	BadActorIdentifierField = "identifier"
)

type BadActor struct {
	Identifier string
	BaId       entity.AutoGenKey
}

func NewBadActor(identifier string) (BadActor, error) {
	var ba BadActor

	identifierTrim := strings.TrimSpace(identifier)
	err := checkIfIdentifierConstraintsAreMet(identifierTrim)

	if err == nil {
		ba = BadActor{
			Identifier: identifierTrim,
		}
	}

	return ba, err
}

func (ba BadActor) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(BadActorIdField, ba.BaId),
		entity.NewTuple(BadActorIdentifierField, ba.Identifier),
	}
}

func (ba BadActor) Copy(key entity.AutoGenKey) BadActor {
	return BadActor{
		BaId:       key,
		Identifier: ba.Identifier,
	}
}

func checkIfIdentifierConstraintsAreMet(identifier string) error {
	size := len(identifier)

	if size == 0 {
		return errors.New("bad actor identifier can not be empty")
	}

	if size > 30 {
		return errors.New("bad actor identifier constraints are not met (max 30 characters)")
	}

	return nil
}
