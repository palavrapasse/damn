package query

type Import struct {
	AffectedUsers     map[User]Credentials
	AffectedPlatforms []Platform
	Leakers           []BadActor
	Leak              Leak
}
