package port

import "fmt"

type settings struct {
	isListening            bool
	uid                    int
	privilegedAllowedPorts []uint16 // allowed ports if running without root
	zeroDisallowed         bool
}

func newSettings(options ...Option) (s *settings, err error) {
	s = new(settings)
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrOption, err)
		}
	}
	return s, nil
}
