package port

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrOption = errors.New("option error")
)

var (
	ErrPortNotAnInteger           = errors.New("port value is not an integer")
	ErrPortLowerThanOne           = errors.New("port cannot be lower than 1")
	ErrPortTooHigh                = errors.New("port cannot be higher than 65535")
	ErrListeningPort              = errors.New("invalid listening port")
	ErrListeningPortLowerThanZero = errors.New("listening port cannot be lower than 0")
)

// Validate verifies the value is a valid port number and
// returns it as an uint16.
func Validate(value string, options ...Option) (port uint16, err error) {
	s, err := newSettings(options...)
	if err != nil {
		return 0, err
	}

	const minPort, maxPort = 1, 65535

	portInt, err := strconv.Atoi(value)
	switch {
	case err != nil:
		return 0, fmt.Errorf("%w: %s", ErrPortNotAnInteger, value)
	case !s.isListening && portInt < minPort:
		return 0, fmt.Errorf("%w: %d", ErrPortLowerThanOne, portInt)
	case s.isListening && portInt < 0:
		return 0, fmt.Errorf("%w: %d", ErrListeningPortLowerThanZero, portInt)
	case portInt > maxPort:
		return 0, fmt.Errorf("%w: %d", ErrPortTooHigh, portInt)
	}

	port = uint16(portInt)

	if s.isListening {
		err = checkListeningPort(port, s.uid, s.privilegedAllowedPorts)
		if err != nil {
			return 0, fmt.Errorf("%w: %w", ErrListeningPort, err)
		}
	}

	return port, nil
}

var (
	errPrivilegedPort = errors.New("cannot use privileged ports (1 to 1023) when running without root")
)

func checkListeningPort(port uint16, uid int,
	allowedPrivilegedPorts []uint16) (err error) {
	const (
		maxPrivilegedPort = 1023
		minDynamicPort    = 49151
	)
	if port == 0 || port > maxPrivilegedPort {
		return nil
	}

	switch uid {
	case -1, 0: // root and windows
		return nil
	default:
		for _, allowed := range allowedPrivilegedPorts {
			if allowed == port {
				return nil
			}
		}
		return fmt.Errorf("%w: %d", errPrivilegedPort, port)
	}
}
