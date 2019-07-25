package streng_test

import (
	"github.com/reiver/go-streng"

	"database/sql/driver"
	"errors"

	"testing"
)

func TestNullableValue(t *testing.T) {

	tests := []struct{
		Nullable      streng.Nullable
		ExpectedValue driver.Value
		ExpectedErr   error
	}{
		{
			Nullable: streng.Nothing().Nullable(),
			ExpectedValue: nil,
			ExpectedErr: errors.New("streng: Nothing"),
		},



		{
			Nullable: streng.Null(),
			ExpectedValue: nil,
			ExpectedErr: nil,
		},



		{
			Nullable: streng.Something("").Nullable(),
			ExpectedValue:             "",
			ExpectedErr: nil,
		},



		{
			Nullable: streng.Something("apple").Nullable(),
			ExpectedValue:             "apple",
			ExpectedErr: nil,
		},
		{
			Nullable: streng.Something("BANANA").Nullable(),
			ExpectedValue:             "BANANA",
			ExpectedErr: nil,
		},
		{
			Nullable: streng.Something("Cherry").Nullable(),
			ExpectedValue:             "Cherry",
			ExpectedErr: nil,
		},
		{
			Nullable: streng.Something("dATE").Nullable(),
			ExpectedValue:             "dATE",
			ExpectedErr: nil,
		},



		{
			Nullable: streng.Something("Hello world!").Nullable(),
			ExpectedValue:             "Hello world!",
			ExpectedErr: nil,
		},



		{
			Nullable: streng.Something("123").Nullable(),
			ExpectedValue:             "123",
			ExpectedErr: nil,
		},
		{
			Nullable: streng.Something("123.45").Nullable(),
			ExpectedValue:             "123.45",
			ExpectedErr: nil,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualErr := test.Nullable.Value()

		if expected, actual := test.ExpectedErr, actualErr; nil == actual && nil != expected {
			t.Errorf("For test #%d, expected error is non-nil, but actual error is nil.", testNumber)
			t.Logf("EXPECTED: (%T) %q", expected, expected)
			t.Logf("ACTUAL:        %#v", actual)
			continue
		}
		if expected, actual := test.ExpectedErr, actualErr; nil != actual && nil == expected {
			t.Errorf("For test #%d, expected error is nil, but actual error is non-nil.", testNumber)
			t.Logf("EXPECTED:    %#v", expected)
			t.Logf("ACTUAL: (%T) %q", actual, actual)
			continue
		}
		if expected, actual := test.ExpectedErr, actualErr; nil != actual && nil != expected {
			if e, a := expected.Error(), actual.Error(); e != a {
				t.Errorf("For test #%d, expected error is not what was actually gotten.", testNumber)
				t.Logf("EXPECTED: (%T) %q", expected, expected)
				t.Logf("ACTUAL: (%T) %q", actual, actual)
				continue
			}
		}

		if expected, actual := test.ExpectedValue, actualValue; expected != actual {
			t.Errorf("For test #%d, what was actually gotten was not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL: %#v", actual)
			t.Logf("EXPECTED error: (%T) %q", test.ExpectedErr, test.ExpectedErr)
			t.Logf("ACTUAL   error: (%T) %q", actualErr, actualErr)
			continue
		}
	}
}
