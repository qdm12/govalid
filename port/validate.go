package port

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrOption           = errors.New("option error")
	ErrPortNotAnInteger = errors.New("port value is not an integer")
	ErrPortNegative     = errors.New("port cannot be negative")
	ErrPortZero         = errors.New("port cannot be zero")
	ErrPortTooHigh      = errors.New("port cannot be higher than 65535")
)

// Validate verifies the value is a valid port number and
// returns it as an uint16.
func Validate(value string, options ...Option) (port uint16, err error) {
	s, err := newSettings(options...)
	if err != nil {
		return 0, err
	}

	const maxPort = 65535

	portInt, err := strconv.Atoi(value)
	switch {
	case err != nil:
		return 0, fmt.Errorf("%w: %s", ErrPortNotAnInteger, value)
	case portInt > maxPort:
		return 0, fmt.Errorf("%w: %d", ErrPortTooHigh, portInt)
	case portInt < 0:
		return 0, fmt.Errorf("%w: %d", ErrPortNegative, portInt)
	}

	port = uint16(portInt)
	if !s.isListening {
		if port == 0 {
			return 0, fmt.Errorf("%w: %d", ErrPortZero, portInt)
		}
		return port, nil
	}

	err = checkListeningPort(port, s.zeroDisallowed, s.uid, s.privilegedAllowedPorts)
	if err != nil {
		return 0, err
	}
	return port, nil
}

var (
	ErrListenPortZero       = errors.New("listening port cannot be zero")
	ErrListenPrivilegedPort = errors.New("listening port cannot be privileged port")
)

func checkListeningPort(port uint16, zeroDisallowed bool,
	uid int, allowedPrivilegedPorts []uint16) (err error) {
	if port == 0 {
		if zeroDisallowed {
			return fmt.Errorf("%w", ErrListenPortZero)
		}
		return nil
	}

	const maxPrivilegedPort = 1023
	if port > maxPrivilegedPort {
		return nil
	}

	if uid == 0 || uid == -1 { // root or windows
		return nil
	}

	if len(allowedPrivilegedPorts) == 0 {
		return fmt.Errorf("%w: %d when running with uid %d", ErrListenPrivilegedPort, port, uid)
	}

	allowedPrivilegedPortsStrings := make([]string, len(allowedPrivilegedPorts))
	for i, allowed := range allowedPrivilegedPorts {
		if allowed == port {
			return nil
		}
		allowedPrivilegedPortsStrings[i] = fmt.Sprint(allowed)
	}

	return fmt.Errorf("%w: port %d is not part of allowed ports %s",
		ErrListenPrivilegedPort, port, strings.Join(allowedPrivilegedPortsStrings, ", "))
}
