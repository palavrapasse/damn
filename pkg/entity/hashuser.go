package entity

type HashUser struct {
	UserId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashUser(us User) HashUser {
	return HashUser{
		UserId:  us.UserId,
		HSHA256: NewHSHA256(string(us.Email)),
	}
}

func (hu HashUser) Record() []Tuple {
	return []Tuple{
		NewTuple("userid", hu.UserId),
		NewTuple("hsha256", hu.HSHA256),
	}
}
