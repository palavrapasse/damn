package query

import (
	"errors"
	"strings"
)

type Password string

func NewPassword(password string) (Password, error) {
	var p Password

	passwordTrim := strings.TrimSpace(password)
	err := checkIfPasswordConstraintsAreMet(passwordTrim)

	if err == nil {
		p = Password(passwordTrim)
	}

	return p, err
}

func checkIfPasswordConstraintsAreMet(p string) error {
	size := len(p)

	if size == 0 {
		return errors.New("password can not be empty")
	}

	return nil
}
