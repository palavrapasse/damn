package query

import (
	"errors"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	CredentialsIdField       = "credid"
	CredentialsPasswordField = "password"
)

type Password string

type Credentials struct {
	Password Password
	CredId   entity.AutoGenKey
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

func (c Credentials) Copy(key entity.AutoGenKey) Credentials {
	return Credentials{
		CredId:   key,
		Password: c.Password,
	}
}

func (c Credentials) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(CredentialsIdField, c.CredId),
		entity.NewTuple(CredentialsPasswordField, c.Password),
	}
}

func checkIfPasswordConstraintsAreMet(p string) error {
	size := len(p)

	if size == 0 {
		return errors.New("password can not be empty")
	}

	return nil
}
