package address

import (
	"github.com/qdm12/govalid/port"
)

type Option func(s *settings) (err error)

func OptionListening(uid int, options ...port.OptionListeningPort) Option {
	return func(s *settings) (err error) {
		portOption := port.OptionPortListening(uid, options...)
		s.portOptions = append(s.portOptions, portOption)
		return nil
	}
}
