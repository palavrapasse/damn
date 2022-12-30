package entity

import (
	"errors"
	"strings"
)

type Password string

type Credentials struct {
	Password Password
	CredId   AutoGenKey
}

func NewPassword(password string) (Password, error) {
	var p Password

	passwordTrim := strings.TrimSpace(password)
	err := checkIfPasswordConstraintsAreMet(passwordTrim)

	if err == nil {
		p = Password(passwordTrim)
	}

	return p, err
}

func NewCredentials(password Password) Credentials {
	return Credentials{
		Password: password,
	}
}

func (c Credentials) Copy(key AutoGenKey) Credentials {
	return Credentials{
		CredId:   key,
		Password: c.Password,
	}
}

func (c Credentials) Record() []Tuple {
	return []Tuple{
		NewTuple("credid", c.CredId),
		NewTuple("password", c.Password),
	}
}

func checkIfPasswordConstraintsAreMet(p string) error {
	size := len(p)

	if size == 0 {
		return errors.New("password can not be empty")
	}

	return nil
}
