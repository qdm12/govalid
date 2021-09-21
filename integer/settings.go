package integer

type settings struct {
	minimum *int
	maximum *int
}

func newSettings() *settings {
	return &settings{}
}
