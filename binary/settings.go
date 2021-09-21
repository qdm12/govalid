package binary

type settings struct {
	enabled  []string
	disabled []string
}

func newSettings() *settings {
	return &settings{
		enabled:  []string{"enabled", "yes", "on"},
		disabled: []string{"disabled", "no", "off"},
	}
}
