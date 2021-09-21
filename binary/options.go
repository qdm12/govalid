package binary

type Option func(s *settings) (err error)

// OptionEnabled sets the values to match to signal
// an enabled result.
func OptionEnabled(values ...string) Option {
	return func(s *settings) (err error) {
		s.enabled = values
		return nil
	}
}

// OptionDisabled sets the values to match to signal
// a disabled result.
func OptionDisabled(values ...string) Option {
	return func(s *settings) (err error) {
		s.disabled = values
		return nil
	}
}
