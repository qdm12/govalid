package port

type settings struct {
	isListening            bool
	uid                    int
	privilegedAllowedPorts []uint16 // allowed ports if running without root
}

func newSettings() *settings {
	return &settings{}
}
