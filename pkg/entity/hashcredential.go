package entity

type HashCredentials struct {
	HSHA256 HSHA256
	CredId  AutoGenKey
}

func NewHashCredentials(cr Credentials) HashCredentials {
	return HashCredentials{
		CredId:  cr.CredId,
		HSHA256: NewHSHA256(string(cr.Password)),
	}
}

func (hc HashCredentials) Record() []Tuple {
	return []Tuple{
		NewTuple(CredentialsIdField, hc.CredId),
		NewTuple(HSHA256IdField, hc.HSHA256),
	}
}
