package errwrap

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrDef interface {
	New(a ...any) ErrWrap
}

type errDefImpl struct {
	code   string
	format string
}

func (e *errDefImpl) New(a ...any) ErrWrap {
	errWrap := &errWrapImpl{
		errDef:    e,
		rawErrStr: fmt.Sprintf(e.format, a...),
		stack:     captureStackTrace(),
	}
	return errWrap
}

func NewErrDef(code string, format string) ErrDef {
	return &errDefImpl{
		code:   code,
		format: format,
	}
}

type ErrWrap interface {
	Error() string
	Code() string
	RawError() string
	StackTrace() string
}

type errWrapImpl struct {
	errDef    *errDefImpl
	rawErrStr string
	stack     string
}

func (e *errWrapImpl) Error() string {
	return fmt.Sprintf("[%v] %v", e.errDef.code, e.rawErrStr)
}

func (e *errWrapImpl) Code() string {
	return e.errDef.code
}

func (e *errWrapImpl) RawError() string {
	return e.rawErrStr
}

func (e *errWrapImpl) StackTrace() string {
	return e.stack
}

func captureStackTrace() string {
	var sb strings.Builder
	pcs := make([]uintptr, 1)
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s:%d", frame.File, frame.Line))
		if !more {
			break
		}
	}

	return sb.String()
}
