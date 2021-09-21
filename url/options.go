package url

type Option func(s *settings) (err error)

// SetAllowedSchemes sets the list of schemes allowed for the URL.
func SetAllowedSchemes(allowedSchemes []string) Option {
	return func(s *settings) (err error) {
		s.allowedSchemes = allowedSchemes
		return nil
	}
}
