package entity

type LeakUser struct {
	UserId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakUser(user User, leak Leak) LeakUser {
	return LeakUser{
		UserId: user.UserId,
		LeakId: leak.LeakId,
	}
}
