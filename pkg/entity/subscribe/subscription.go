package subscribe

type Subscription struct {
	Subscriber Subscriber
	Affected   []Affected
}
