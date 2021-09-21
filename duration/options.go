package duration

type Option func(s *settings) (err error)

// OptionAllowNegative allows the parsed duration to be negative.
func OptionAllowNegative() Option {
	return func(s *settings) (err error) {
		s.allowNegative = true
		return nil
	}
}

// OptionAllowZero allows the parsed duration to be zero.
func OptionAllowZero() Option {
	return func(s *settings) (err error) {
		s.allowZero = true
		return nil
	}
}
