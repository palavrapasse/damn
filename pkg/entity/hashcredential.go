package entity

type HashCredentials struct {
	CredId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashCredentials(cr Credentials) HashCredentials {
	return HashCredentials{
		CredId:  cr.CredId,
		HSHA256: NewHSHA256(string(cr.Password)),
	}
}

func (hc HashCredentials) Record() []Tuple {
	return []Tuple{
		NewTuple("credid", hc.CredId),
		NewTuple("hsha256", hc.HSHA256),
	}
}
