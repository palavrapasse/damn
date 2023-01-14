package entity

const (
	SubscriberIdField       = "subid"
	SubscriberB64EmailField = "b64email"
)

type Subscriber struct {
	B64Email     string
	SubscriberId AutoGenKey
}

func NewSubscriber(emailB64 string) Subscriber {
	return Subscriber{
		B64Email: emailB64,
	}
}

func (s Subscriber) Copy(key AutoGenKey) Subscriber {
	return Subscriber{
		SubscriberId: key,
		B64Email:     s.B64Email,
	}
}

func (s Subscriber) Record() []Tuple {
	return []Tuple{
		NewTuple(AffectedIdField, s.SubscriberId),
		NewTuple(AffectedHashEmailField, s.B64Email),
	}
}
