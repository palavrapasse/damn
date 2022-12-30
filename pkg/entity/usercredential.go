package entity

type UserCredentials struct {
	CredId AutoGenKey
	UserId AutoGenKey
}

func NewUserCredentials(cred Credentials, user User) UserCredentials {
	return UserCredentials{
		CredId: cred.CredId,
		UserId: user.UserId,
	}
}
