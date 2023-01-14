package query

import "github.com/palavrapasse/damn/pkg/entity"

type LeakBadActor struct {
	BaId   entity.AutoGenKey
	LeakId entity.AutoGenKey
}

func NewLeakBadActor(ba BadActor, leak Leak) LeakBadActor {
	return LeakBadActor{
		BaId:   ba.BaId,
		LeakId: leak.LeakId,
	}
}

func (lba LeakBadActor) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(BadActorIdField, lba.BaId),
		entity.NewTuple(LeakIdField, lba.LeakId),
	}
}
