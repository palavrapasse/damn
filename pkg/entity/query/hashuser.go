package query

import "github.com/palavrapasse/damn/pkg/entity"

type HashUser struct {
	HSHA256 entity.HSHA256
	UserId  entity.AutoGenKey
}

func NewHashUser(us User) HashUser {
	return HashUser{
		UserId:  us.UserId,
		HSHA256: entity.NewHSHA256(string(us.Email)),
	}
}

func (hu HashUser) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(UserIdField, hu.UserId),
		entity.NewTuple(entity.HSHA256IdField, hu.HSHA256),
	}
}
