package entity

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
