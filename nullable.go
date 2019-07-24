package streng

import (
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

func (receiver Nullable) GoString() string {
	if Nothing().Nullable() == receiver {
		return "streng.Nothing().Nullable()"
	}

	if Null() == receiver {
		return "streng.Null()"
	}

	return fmt.Sprintf("streng.Something(%q).Nullable()", receiver.value)
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
