package entity

type HashUser struct {
	HSHA256 HSHA256
	UserId  AutoGenKey
}

func NewHashUser(us User) HashUser {
	return HashUser{
		UserId:  us.UserId,
		HSHA256: NewHSHA256(string(us.Email)),
	}
}

func (hu HashUser) Record() []Tuple {
	return []Tuple{
		NewTuple(UserIdField, hu.UserId),
		NewTuple(HSHA256IdField, hu.HSHA256),
	}
}
