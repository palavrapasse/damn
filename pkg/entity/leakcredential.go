package entity

type LeakCredentials struct {
	CredId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakCredentials(cred Credentials, leak Leak) LeakCredentials {
	return LeakCredentials{
		CredId: cred.CredId,
		LeakId: leak.LeakId,
	}
}

func (lc LeakCredentials) Record() []Tuple {
	return []Tuple{
		NewTuple("credid", lc.CredId),
		NewTuple("leakid", lc.LeakId),
	}
}
