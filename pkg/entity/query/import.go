package query

type Import struct {
	AffectedUsers     []User
	AffectedPlatforms []Platform
	Leakers           []BadActor
	Leak              Leak
}
