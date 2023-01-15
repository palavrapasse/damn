package query

import "github.com/palavrapasse/damn/pkg/entity"

type LeakPlatform struct {
	PlatId entity.AutoGenKey
	LeakId entity.AutoGenKey
}

func NewLeakPlatform(plat Platform, leak Leak) LeakPlatform {
	return LeakPlatform{
		PlatId: plat.PlatId,
		LeakId: leak.LeakId,
	}
}

func (lpt LeakPlatform) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(PlatformIdField, lpt.PlatId),
		entity.NewTuple(LeakIdField, lpt.LeakId),
	}
}
