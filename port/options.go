package port

import "fmt"

type Option func(s *settings) (err error)
type OptionListeningPort func(s *settings) (err error)

func OptionPortListening(uid int, options ...OptionListeningPort) Option {
	return func(s *settings) (err error) {
		s.isListening = true
		s.uid = uid

		for _, option := range options {
			err := option(s)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrOption, err)
			}
		}

		return nil
	}
}

// OptionListeningPortPrivilegedAllowed sets a list of privileged allowed ports
// if listening as non-root.
func OptionListeningPortPrivilegedAllowed(ports ...uint16) OptionListeningPort {
	return func(s *settings) (err error) {
		s.privilegedAllowedPorts = ports
		return nil
	}
}
