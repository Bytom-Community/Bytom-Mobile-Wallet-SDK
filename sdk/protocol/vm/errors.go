package vm

import "errors"

var (
	ErrAltStackUnderflow  = errors.New("alt stack underflow")
	ErrBadValue           = errors.New("bad value")
	ErrContext            = errors.New("wrong context")
	ErrDataStackUnderflow = errors.New("data stack underflow")
	ErrDisallowedOpcode   = errors.New("disallowed opcode")
	ErrDivZero            = errors.New("division by zero")
	ErrFalseVMResult      = errors.New("false VM result")
	ErrLongProgram        = errors.New("program size exceeds maxint32")
	ErrRange              = errors.New("range error")
	ErrReturn             = errors.New("RETURN executed")
	ErrRunLimitExceeded   = errors.New("run limit exceeded")
	ErrShortProgram       = errors.New("unexpected end of program")
	ErrToken              = errors.New("unrecognized token")
	ErrUnexpected         = errors.New("unexpected error")
	ErrUnsupportedVM      = errors.New("unsupported VM")
	ErrVerifyFailed       = errors.New("VERIFY failed")
)
