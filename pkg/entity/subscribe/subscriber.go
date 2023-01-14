package subscribe

import "github.com/palavrapasse/damn/pkg/entity"

const (
	SubscriberIdField       = "subid"
	SubscriberB64EmailField = "b64email"
)

type Subscriber struct {
	B64Email     entity.Base64
	SubscriberId entity.AutoGenKey
}

func NewSubscriber(email string) Subscriber {
	return Subscriber{
		B64Email: entity.NewBase64(email),
	}
}

func (s Subscriber) Copy(key entity.AutoGenKey) Subscriber {
	return Subscriber{
		SubscriberId: key,
		B64Email:     s.B64Email,
	}
}

func (s Subscriber) Record() []entity.Tuple {
	return []entity.Tuple{
		entity.NewTuple(SubscriberIdField, s.SubscriberId),
		entity.NewTuple(SubscriberB64EmailField, s.B64Email),
	}
}
