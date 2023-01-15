package query

import (
	"errors"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	LeakIdField          = "leakid"
	LeakContextField     = "context"
	LeakShareDateSCField = "sharedatesc"
)

type Context string

type Leak struct {
	Context     Context
	ShareDateSC DateInSeconds
	LeakId      entity.AutoGenKey
}

func NewLeak(context string, shareDateSC DateInSeconds) (Leak, error) {
	var l Leak

	contextTrim := strings.TrimSpace(context)
	err := checkIfContextConstraintsAreMet(contextTrim)

	if err == nil {
		l = Leak{
			Context:     Context(contextTrim),
			ShareDateSC: shareDateSC,
		}
	}

	return l, err
}

func (l Leak) Copy(key entity.AutoGenKey) Leak {
	return Leak{
		LeakId:      key,
		ShareDateSC: l.ShareDateSC,
		Context:     l.Context,
	}
}

func (l Leak) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(LeakIdField, l.LeakId),
		entity.NewTuple(LeakShareDateSCField, l.ShareDateSC),
		entity.NewTuple(LeakContextField, l.Context),
	}
}

func checkIfContextConstraintsAreMet(c string) error {
	size := len(c)

	if size == 0 {
		return errors.New("leak context can not be empty")
	}

	if size > 130 {
		return errors.New("leak context constraints are not met (max 130 characters)")
	}

	return nil
}
