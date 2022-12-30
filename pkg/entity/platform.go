package entity

import (
	"errors"
	"strings"
)

type Platform struct {
	PlatId AutoGenKey
	Name   string
}

func NewPlatform(name string) (Platform, error) {
	var p Platform

	nameTrim := strings.TrimSpace(name)
	err := checkIfPlatformNameConstraintsAreMet(nameTrim)

	if err == nil {
		p = Platform{
			Name: nameTrim,
		}
	}

	return p, err
}

func (p Platform) Copy(key AutoGenKey) Platform {
	return Platform{
		PlatId: key,
		Name:   p.Name,
	}
}

func checkIfPlatformNameConstraintsAreMet(n string) error {
	size := len(n)

	if size == 0 {
		return errors.New("platform name can not be empty")
	}

	if size > 30 {
		return errors.New("platform name constraints are not met (max 30 characters)")
	}

	return nil
}
