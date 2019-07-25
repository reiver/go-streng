package streng

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Nullable struct {
	value  string
	isnull bool
	loaded bool
}

func Null() Nullable {
	return Nullable{
		loaded: true,
		isnull: true,
	}
}

func someNullable(value string) Nullable {
	return Nullable{
		loaded: true,
		value:  value,
	}
}

// Else defaults this ‘streng.Nullable’ to ‘value’ if this ‘streng.Nullable’ has a value of ‘streng.Nothing().Nullable()’,
// or a value of ‘streng.Null()’, else it just returns itself as is.
func (receiver Nullable) Else(value string) Nullable {
        if Nothing().Nullable() == receiver {
                return Something(value).Nullable()
        }
        if Null() == receiver {
                return Something(value).Nullable()
        }

        return receiver
}
func (receiver Nullable) GoString() string {
	if Nothing().Nullable() == receiver {
		return "streng.Nothing().Nullable()"
	}
	if Null() == receiver {
		return "streng.Null()"
	}

	return fmt.Sprintf("streng.Something(%q).Nullable()", receiver.value)
}

// Map returns an ‘streng.Nullable’ containing the result of ‘fn’ applied to the value inside this ‘streng.Nullable’;
// if this ‘streng.Nullable’ is ‘streng.Nothing().Nullable()’, then it just returns ‘streng.Nothing()’;
// or if this ‘streng.Nullable’ is ‘streng.Null()’, then it just returns ‘streng.Null()’;
func (receiver Nullable) Map(fn func(string)string) Nullable {
	if Nothing().Nullable() == receiver {
		return receiver
	}
	if Null() == receiver {
		return receiver
	}

	return Something(fn(receiver.value)).Nullable()
}

func (receiver Nullable) MarshalJSON() ([]byte, error) {
	if Nothing().Nullable() == receiver {
		return nil, errNothing
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver Nullable) Return() (string, error) {
        if Nothing().Nullable() == receiver {
                return "", errNothing
        }
        if Null() == receiver {
                return "", errNull
        }

        return receiver.value, nil
}

// Then returns an ‘streng.Nullable’ containing the result of ‘fn’ applied to the value inside this ‘streng.Nullable’;
// if this ‘streng.Nullable’ is ‘streng.Nothing().Nullable()’, then it just returns ‘streng.Nothing()’;
// and if this ‘streng.Nullable’ is ‘streng.Null()’, then it just returns ‘streng.Null()’.
func (receiver Nullable) Then(fn func(string)Nullable) Nullable {
	if Nothing() == receiver {
		return receiver
	}

	return fn(receiver.value)
}

// UnmarshalJSON makes streng.Nullable fit the json.Unmarshaler interface.
//
// Note that streng.Nullable.UnmarshalJSON() also accepts JSON numbers, and not just JSON strings.
//
// So if we have:
//
//	type Purchase struct {
//		Amount streng.Nullable `json:"amount"`
//	}
//
// It accepts this JSON with a string literal:
//
//	{
//		"amount": "1.23"
//	}
//
// But it also accepts this JSON with a number literal:
//
//	{
//		"amount": 1.23
//	}
func (receiver *Nullable) UnmarshalJSON(data []byte) error {
	if nil == receiver {

		return errNilReceiver
	}
	if nil == data {
		return fmt.Errorf("streng: %#v is invalid JSON", data)
	}

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}
	var str string = *(*string)(unsafe.Pointer(&stringHeader))

	switch {
	case "null" == str:
		*receiver = Null()
		return nil
	case 2 <= len(str) && strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`):
		var dest string
		if err := json.Unmarshal(data, &dest); nil != err {
			return err
		}

		*receiver = someNullable(dest)
		return nil
	default:
		*receiver = someNullable(string(data))
		return nil
	}
}

func (receiver Nullable) Unwrap() (string, bool) {
	if  Nothing().Nullable() == receiver {
		return "", false
	}
	if Null() == receiver {
		return "", false
	}

	return receiver.value, true
}

// Value makes ‘streng.Nullable’ fit the database/sql/driver.Valuer interface.
func (receiver Nullable) Value() (driver.Value, error) {
	if Nothing().Nullable() == receiver {
		return receiver, errNothing
	}
	if Null() == receiver {
		return nil, nil
	}

	return receiver.value, nil
}
