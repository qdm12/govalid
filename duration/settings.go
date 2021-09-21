package duration

type settings struct {
	allowNegative bool
	allowZero     bool
}

func newSettings() *settings {
	return &settings{}
}
