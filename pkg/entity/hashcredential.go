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
