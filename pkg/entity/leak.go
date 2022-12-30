package entity

import (
	"errors"
	"strings"
)

type Context string

type Leak struct {
	Context     Context
	ShareDateSC DateInSeconds
	LeakId      AutoGenKey
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

func (l Leak) Copy(key AutoGenKey) Leak {
	return Leak{
		LeakId:      key,
		ShareDateSC: l.ShareDateSC,
		Context:     l.Context,
	}
}

func (l Leak) Record() []Tuple {
	return []Tuple{
		NewTuple("leakid", l.LeakId),
		NewTuple("sharedatesc", l.ShareDateSC),
		NewTuple("context", l.Context),
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
