package entity

type LeakBadActor struct {
	BaId   AutoGenKey
	LeakId AutoGenKey
}

func NewLeakBadActor(ba BadActor, leak Leak) LeakBadActor {
	return LeakBadActor{
		BaId:   ba.BaId,
		LeakId: leak.LeakId,
	}
}
