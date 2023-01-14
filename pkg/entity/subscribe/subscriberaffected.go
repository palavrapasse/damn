package subscribe

import "github.com/palavrapasse/damn/pkg/entity"

type SubscriberAffected struct {
	AffId entity.AutoGenKey
	SubId entity.AutoGenKey
}

func NewSubscriberAffected(a Affected, s Subscriber) SubscriberAffected {
	return SubscriberAffected{
		AffId: a.AffectedId,
		SubId: s.SubscriberId,
	}
}

func (sa SubscriberAffected) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(AffectedIdField, sa.AffId),
		entity.NewTuple(SubscriberIdField, sa.SubId),
	}
}
