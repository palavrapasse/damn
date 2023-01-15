package query

import "github.com/palavrapasse/damn/pkg/entity"

type LeakCredentials struct {
	CredId entity.AutoGenKey
	LeakId entity.AutoGenKey
}

func NewLeakCredentials(cred Credentials, leak Leak) LeakCredentials {
	return LeakCredentials{
		CredId: cred.CredId,
		LeakId: leak.LeakId,
	}
}

func (lc LeakCredentials) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(CredentialsIdField, lc.CredId),
		entity.NewTuple(LeakIdField, lc.LeakId),
	}
}
