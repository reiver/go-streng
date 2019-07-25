package streng_test

import (
	"github.com/reiver/go-streng"

	"database/sql/driver"
	"errors"

	"testing"
)

func TestOptionValue(t *testing.T) {

	tests := []struct{
		Option        streng.Option
		ExpectedValue driver.Value
		ExpectedErr   error
	}{
		{
			Option: streng.Nothing(),
			ExpectedValue: nil,
			ExpectedErr: errors.New("streng: Nothing"),
		},



		{
			Option: streng.Something(""),
			ExpectedValue:           "",
			ExpectedErr: nil,
		},



		{
			Option: streng.Something("apple"),
			ExpectedValue:           "apple",
			ExpectedErr: nil,
		},
		{
			Option: streng.Something("BANANA"),
			ExpectedValue:           "BANANA",
			ExpectedErr: nil,
		},
		{
			Option: streng.Something("Cherry"),
			ExpectedValue:           "Cherry",
			ExpectedErr: nil,
		},
		{
			Option: streng.Something("dATE"),
			ExpectedValue:           "dATE",
			ExpectedErr: nil,
		},



		{
			Option: streng.Something("Hello world!"),
			ExpectedValue:           "Hello world!",
			ExpectedErr: nil,
		},



		{
			Option: streng.Something("123"),
			ExpectedValue:           "123",
			ExpectedErr: nil,
		},
		{
			Option: streng.Something("123.45"),
			ExpectedValue:           "123.45",
			ExpectedErr: nil,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualErr := test.Option.Value()

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
