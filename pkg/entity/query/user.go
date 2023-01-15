package query

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	UserIdField    = "userid"
	UserEmailField = "email"
)

type Email string

type User struct {
	Email  Email
	UserId entity.AutoGenKey
}

func NewUser(email string) (User, error) {
	var u User

	emailTrim := strings.TrimSpace(email)
	err := checkIfEmailConstraintsAreMet(emailTrim)

	if err == nil {
		u = User{
			Email: Email(emailTrim),
		}
	}

	return u, err
}

func (u User) Copy(key entity.AutoGenKey) User {
	return User{
		UserId: key,
		Email:  u.Email,
	}
}

func (u User) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(UserIdField, u.UserId),
		entity.NewTuple(UserEmailField, u.Email),
	}
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
