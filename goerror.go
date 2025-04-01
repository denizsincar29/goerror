// simple error handling package for golang
package goerror

import (
	"github.com/sirupsen/logrus"
)

// Error struct
type Error struct {
	// last error
	lastError error
	// logrus logger
	logger *logrus.Logger
}

// NewError creates a new Error struct
func NewError(logger *logrus.Logger) *Error {
	e := new(Error)
	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.DebugLevel)
	}
	e.logger = logger
	return e
}

// Must takes a error and panics if the error is not nil
func (e *Error) Must(err error, args ...string) {
	if err != nil {
		if len(args) > 0 {
			e.logger.Fatalln(args[0], err.Error())
		} else {
			e.logger.Fatalln(err.Error())
		}
	}
}

// Check takes a error and logs it if it is not nil
func (e *Error) Check(err error, args ...string) {
	e.lastError = err
	if err != nil {
		if len(args) > 0 {
			e.logger.Errorln(args[0], err.Error())
		} else {
			e.logger.Errorln(err.Error())
		}
	}
}

// warn takes a error and logs it if it is not nil
func (e *Error) Warn(err error, args ...string) {
	e.lastError = err
	if err != nil {
		if len(args) > 0 {
			e.logger.Warnln(args[0], err.Error())
		} else {
			e.logger.Warnln(err.Error())
		}
	}
}

// Info takes a error and logs it if it is not nil
func (e *Error) Info(err error, args ...string) {
	e.lastError = err
	// this function is used in rare cases
	if err != nil {
		if len(args) > 0 {
			e.logger.Infoln(args[0], err.Error())
		} else {
			e.logger.Infoln(err.Error())
		}
	}
}

// Debug takes a error and logs it if it is not nil
func (e *Error) Debug(err error, args ...string) {
	e.lastError = err
	// this function is used in rare cases
	if err != nil {
		if len(args) > 0 {
			e.logger.Debugln(args[0], err.Error())
		} else {
			e.logger.Debugln(err.Error())
		}
	}
}

// IsError returns true if the last error is not nil
func (e *Error) IsError() bool {
	return e.lastError != nil
}

// E alias for iserror
func (e *Error) E() bool {
	return e.IsError()
}
