package url

type settings struct {
	allowedSchemes []string
}

func newSettings() *settings {
	return &settings{
		allowedSchemes: []string{"http", "https"},
	}
}
