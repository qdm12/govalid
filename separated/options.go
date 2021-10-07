package separated

type Option func(s *settings) (err error)

// OptionSeparator sets the separator to use to
// split the value string.
func OptionSeparator(separator string) Option {
	return func(s *settings) (err error) {
		s.separator = separator
		return nil
	}
}

// OptionAccepted sets a list of accepted values for each
// separated value.
func OptionAccepted(accepted ...string) Option {
	return func(s *settings) (err error) {
		s.accepted = accepted
		return nil
	}
}

// OptionLowercase changes all the values to lowercase.
func OptionLowercase() Option {
	return func(s *settings) (err error) {
		s.lowercase = true
		return nil
	}
}

// OptionIgnoreEmpty ignores the empty string value from split values.
func OptionIgnoreEmpty() Option {
	return func(s *settings) (err error) {
		s.ignoreEmpty = true
		return nil
	}
}
