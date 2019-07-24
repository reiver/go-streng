package streng

import (
	"errors"
)

var (
	errNilReceiver = errors.New("streng: Nil Receiver")
	errNothing     = errors.New("streng: Nothing")
	errNull        = errors.New("streng: Null")
)
