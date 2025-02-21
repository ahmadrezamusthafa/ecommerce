package apperror

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Error struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Cause   error  `json:"cause"`
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s:%d: %s: %v", e.File, e.Line, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func New(message string) error {
	file, line := getCallerInfo()
	return &Error{
		Message: message,
		File:    file,
		Line:    line,
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	file, line := getCallerInfo()
	return &Error{
		Message: message,
		File:    file,
		Line:    line,
		Cause:   err,
	}
}

func getCallerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", 0
	}
	return filepath.Base(file), line
}
