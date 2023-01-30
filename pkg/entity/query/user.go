package query

import (
	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	UserIdField    = "userid"
	UserEmailField = "email"
)

type User struct {
	Email  Email
	UserId entity.AutoGenKey
}

func NewUser(email Email) User {
	return User{
		Email: email,
	}
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
