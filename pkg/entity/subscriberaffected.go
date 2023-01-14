package entity

type SubscriberAffected struct {
	AffId AutoGenKey
	SubId AutoGenKey
}

func NewSubscriberAffected(cred Credentials, user User) SubscriberAffected {
	return SubscriberAffected{
		AffId: cred.CredId,
		SubId: user.UserId,
	}
}

func (sa SubscriberAffected) Record() []Tuple {
	return []Tuple{
		NewTuple(CredentialsIdField, sa.AffId),
		NewTuple(UserIdField, sa.SubId),
	}
}
