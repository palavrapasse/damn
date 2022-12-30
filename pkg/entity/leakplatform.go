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
