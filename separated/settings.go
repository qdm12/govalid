package separated

import "fmt"

type settings struct {
	separator   string
	accepted    []string
	lowercase   bool
	ignoreEmpty bool
}

func newSettings(options ...Option) (s *settings, err error) {
	s = &settings{
		separator: ",",
	}
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrOption, err)
		}
	}
	return s, nil
}
