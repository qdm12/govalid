package url

type Option func(s *settings) (err error)

// OptionAllowSchemes sets the list of schemes allowed for the URL.
func OptionAllowSchemes(allowedSchemes ...string) Option {
	return func(s *settings) (err error) {
		s.allowedSchemes = allowedSchemes
		return nil
	}
}
