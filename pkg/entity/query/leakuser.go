package query

import "github.com/palavrapasse/damn/pkg/entity"

type LeakUser struct {
	UserId entity.AutoGenKey
	LeakId entity.AutoGenKey
}

func NewLeakUser(user User, leak Leak) LeakUser {
	return LeakUser{
		UserId: user.UserId,
		LeakId: leak.LeakId,
	}
}

func (lu LeakUser) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(UserIdField, lu.UserId),
		entity.NewTuple(LeakIdField, lu.LeakId),
	}
}
