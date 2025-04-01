// simple error handling package for golang
package goerror

import (
	"fmt"
	"log/slog"
	"os"
)

// Error struct
type Error struct {
	// last error
	lastError error
	// slogger
	logger *slog.Logger
}

// NewError creates a new Error struct
func NewError(logger *slog.Logger) *Error {
	e := new(Error)
	if logger == nil {
		// create a new logger with default settings
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		logger.Info("Don't do this! Use your own logger")
	}
	e.logger = logger
	return e
}

// Must takes a error and panics if the error is not nil
func (e *Error) Must(err error, args ...string) {
	if err != nil {
		if len(args) > 0 {
			panic(fmt.Sprintf("%s: %s", args[0], err.Error()))
		} else {
			panic(err.Error())
		}
	}
}

// Check takes a error and logs it if it is not nil
func (e *Error) Check(err error, args ...string) {
	e.lastError = err
	if err != nil {
		if len(args) > 0 {
			e.logger.Error(args[0], err.Error())
		} else {
			e.logger.Error(err.Error())
		}
	}
}

// warn takes a error and logs it if it is not nil
func (e *Error) Warn(err error, args ...string) {
	e.lastError = err
	if err != nil {
		if len(args) > 0 {
			e.logger.Warn(args[0], err.Error())
		} else {
			e.logger.Warn(err.Error())
		}
	}
}

// Info takes a error and logs it if it is not nil
func (e *Error) Info(err error, args ...string) {
	e.lastError = err
	// this function is used in rare cases
	if err != nil {
		if len(args) > 0 {
			e.logger.Info(args[0], err.Error())
		} else {
			e.logger.Info(err.Error())
		}
	}
}

// Debug takes a error and logs it if it is not nil
func (e *Error) Debug(err error, args ...string) {
	e.lastError = err
	// this function is used in rare cases
	if err != nil {
		if len(args) > 0 {
			e.logger.Debug(args[0], err.Error())
		} else {
			e.logger.Debug(err.Error())
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
