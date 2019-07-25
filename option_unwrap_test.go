package streng_test

import (
	"github.com/reiver/go-streng"

	"testing"
)

func TestOptionUnwrap(t *testing.T) {

	tests := []struct{
		Option        streng.Option
		ExpectedValue string
		ExpectedOK    bool
	}{
		{
			Option: streng.Nothing(),
			ExpectedValue: "",
			ExpectedOK: false,
		},



		{
			Option: streng.Something(""),
			ExpectedValue:           "",
			ExpectedOK: true,
		},



		{
			Option: streng.Something("apple"),
			ExpectedValue:           "apple",
			ExpectedOK: true,
		},
		{
			Option: streng.Something("BANANA"),
			ExpectedValue:           "BANANA",
			ExpectedOK: true,
		},
		{
			Option: streng.Something("Cherry"),
			ExpectedValue:           "Cherry",
			ExpectedOK: true,
		},
		{
			Option: streng.Something("dATE"),
			ExpectedValue:           "dATE",
			ExpectedOK: true,
		},



		{
			Option: streng.Something("Hello world!"),
			ExpectedValue:           "Hello world!",
			ExpectedOK: true,
		},



		{
			Option: streng.Something("123"),
			ExpectedValue:           "123",
			ExpectedOK: true,
		},
		{
			Option: streng.Something("123.45"),
			ExpectedValue:           "123.45",
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualOK := test.Option.Unwrap()

		if expected, actual := test.ExpectedOK, actualOK; expected != actual {
			t.Errorf("For test #%d, for the OK, that was actually gotten was not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}

		if expected, actual := test.ExpectedValue, actualValue; expected != actual {
			t.Errorf("For test #%d, what was actually gotten was not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL: %#v", actual)
			t.Logf("EXPECTED ok: %t", test.ExpectedOK)
			t.Logf("ACTUAL   ok: %t", actualOK)
			continue
		}
	}
}
