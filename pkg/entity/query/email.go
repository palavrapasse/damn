package query

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

func NewEmail(email string) (Email, error) {
	var e Email

	emailTrim := strings.ToLower(strings.TrimSpace(email))
	err := checkIfEmailConstraintsAreMet(emailTrim)

	if err == nil {
		e = Email(emailTrim)
	}

	return e, err
}

func checkIfEmailConstraintsAreMet(e string) error {
	_, err := mail.ParseAddress(e)

	if err != nil {
		return err
	}

	if len(e) > 130 {
		return errors.New("user email constraints are not met (max 130 characters)")
	}

	return nil
}
