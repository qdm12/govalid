package separated

type settings struct {
	separator   string
	accepted    []string
	lowercase   bool
	ignoreEmpty bool
}

func newSettings() *settings {
	return &settings{
		separator: ",",
	}
}
