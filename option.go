package streng

import (
	"database/sql/driver"
	"fmt"
)

// Option is an ‘option type’ for string.
//
// It can contain:
//
// • ‘nothing’, or
//
// • ‘something’.
//
//
// Nothing
//
// An uninitialize variable, of type ‘streng.Option’, contains ‘nothing’. I.e.,:...
//
//	var option streng.Option
//
// You can determine if a variable of type ‘streng.Option’ contains ‘nothing’, or not,
// by using the ‘streng.Nothing()’ function.
//
// For example:
//
//	if streng.Nothing() == option {
//		//@TODO
//	}
//
// Or:
//
//	switch option {
//	case streng.Nothing():
//		//@TODO
//	default:
//		//@TODO
//	}
//
// You can also use the ‘streng.Nothing()’ function to give a variable of type
// ‘streng.Option’ the value of nothing ‘nothing’:
//
//	option = streng.Nothing()
//
//
// Something
//
// You can create a ‘streng.Option’ with ‘streng.Something()’. For example:
//
//	var option streng.Option = streng.Something("Hello world!")
type Option struct {
	value  string
	loaded bool
}

// Nothing returns an empty streng.Option.
func Nothing() Option {
	return Option{}
}

// Something returns a streng.Option with ‘value’ in it.
func Something(value string) Option {
	return Option{
		loaded: true,
		value:  value,
	}
}

func (receiver Option) GoString() string {
	if Nothing() == receiver {
		return "streng.Nothing()"
	}

	return fmt.Sprintf("streng.Something(%q)", receiver.value)
}

// Else defaults this ‘streng.Option’ to ‘value’ if this ‘streng.Option’ has a value of ‘streng.Nothing()’,
// else it just returns itself as is.
func (receiver Option) Else(value string) Option {
	if Nothing() == receiver {
		return Something(value)
	}

	return receiver
}

func (receiver Option) ElseUnwrap(value string) string {
	if Nothing() == receiver {
		return value
	}

	return receiver.value
}

// Map returns an ‘streng.Option’ containing the result of ‘fn’ applied to the value inside this ‘streng.Option’;
// if this ‘streng.Option’ is ‘streng.Nothing()’, then it just returns ‘streng.Nothing()’.
func (receiver Option) Map(fn func(string)string) Option {
	if Nothing() == receiver {
		return receiver
	}

	return Something(fn(receiver.value))
}

func (receiver Option) MarshalJSON() ([]byte, error) {
	return receiver.Nullable().MarshalJSON()
}

// Nullable returns the equivalent ‘streng.Nullable’ for this ‘streng.Option’.
func (receiver Option) Nullable() Nullable {
	if Nothing() == receiver {
		var nothing Nullable
		return nothing
	}

	return someNullable(receiver.value)
}

// Return returns the string inside, if there is one inside.
func (receiver Option) Return() (string, error) {
	if Nothing() == receiver {
		return "", errNothing
	}

	return receiver.value, nil
}

// Then returns an ‘streng.Option’ containing the result of ‘fn’ applied to the value inside this ‘streng.Option’;
// if this ‘streng.Option’ is ‘streng.Nothing()’, then it just returns ‘streng.Nothing()’.
func (receiver Option) Then(fn func(string)Option) Option {
	if Nothing() == receiver {
		return receiver
	}

	return fn(receiver.value)
}

// Unwrap returns the string inside, if there is one inside.
func (receiver Option) Unwrap() (string, bool) {
	if Nothing() == receiver {
		return "", false
	}

	return receiver.value, true
}

// Value makes ‘streng.Option’ fit the database/sql/driver.Valuer interface.
func (receiver Option) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return receiver, errNothing
	}

	return receiver.value, nil
}
