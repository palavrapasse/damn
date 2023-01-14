package query

import "github.com/palavrapasse/damn/pkg/entity"

type HashCredentials struct {
	HSHA256 entity.HSHA256
	CredId  entity.AutoGenKey
}

func NewHashCredentials(cr Credentials) HashCredentials {
	return HashCredentials{
		CredId:  cr.CredId,
		HSHA256: entity.NewHSHA256(string(cr.Password)),
	}
}

func (hc HashCredentials) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(CredentialsIdField, hc.CredId),
		entity.NewTuple(entity.HSHA256IdField, hc.HSHA256),
	}
}
