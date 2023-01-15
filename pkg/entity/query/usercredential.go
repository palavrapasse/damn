package query

import (
	"github.com/palavrapasse/damn/pkg/entity"
)

type UserCredentials struct {
	CredId entity.AutoGenKey
	UserId entity.AutoGenKey
}

func NewUserCredentials(cred Credentials, user User) UserCredentials {
	return UserCredentials{
		CredId: cred.CredId,
		UserId: user.UserId,
	}
}

func (uc UserCredentials) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(CredentialsIdField, uc.CredId),
		entity.NewTuple(UserIdField, uc.UserId),
	}
}
