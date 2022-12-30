package entity

type Import struct {
	Leak              Leak
	AffectedUsers     map[User]Credentials
	AffectedPlatforms []Platform
	Leakers           []BadActor
}
