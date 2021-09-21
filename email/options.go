package email

type Option func(s *settings) (err error)

// OptionMXLookup instructs the Email function to check
// the corresponding email host does have an MX record.
func OptionMXLookup() Option {
	return func(s *settings) (err error) {
		s.mxLookup = true
		return nil
	}
}
