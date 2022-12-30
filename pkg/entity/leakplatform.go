package entity

type LeakPlatform struct {
	PlatId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakPlatform(plat Platform, leak Leak) LeakPlatform {
	return LeakPlatform{
		PlatId: plat.PlatId,
		LeakId: leak.LeakId,
	}
}

func (lpt LeakPlatform) Record() []Tuple {
	return []Tuple{
		NewTuple(PlatformIdField, lpt.PlatId),
		NewTuple(LeakIdField, lpt.LeakId),
	}
}
