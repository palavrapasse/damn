package entity

import (
	"errors"
	"strings"
)

type BadActor struct {
	BaId       AutoGenKey
	Identifier string
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

func (ba BadActor) Copy(key AutoGenKey) BadActor {
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
