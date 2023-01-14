package subscribe

import "github.com/palavrapasse/damn/pkg/entity"

const (
	SubscriberIdField       = "subid"
	SubscriberB64EmailField = "b64email"
)

type Subscriber struct {
	B64Email     string
	SubscriberId entity.AutoGenKey
}

func NewSubscriber(emailB64 string) Subscriber {
	return Subscriber{
		B64Email: emailB64,
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
		entity.NewTuple(AffectedIdField, s.SubscriberId),
		entity.NewTuple(AffectedHashEmailField, s.B64Email),
	}
}
