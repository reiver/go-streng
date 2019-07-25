package streng_test

import (
	"github.com/reiver/go-streng"

	"testing"
)

func TestNullableUnwrap(t *testing.T) {

	tests := []struct{
		Nullable      streng.Nullable
		ExpectedValue string
		ExpectedOK    bool
	}{
		{
			Nullable: streng.Nothing().Nullable(),
			ExpectedValue: "",
			ExpectedOK: false,
		},



		{
			Nullable: streng.Null(),
			ExpectedValue: "",
			ExpectedOK: false,
		},



		{
			Nullable: streng.Something("").Nullable(),
			ExpectedValue:             "",
			ExpectedOK: true,
		},



		{
			Nullable: streng.Something("apple").Nullable(),
			ExpectedValue:             "apple",
			ExpectedOK: true,
		},
		{
			Nullable: streng.Something("BANANA").Nullable(),
			ExpectedValue:             "BANANA",
			ExpectedOK: true,
		},
		{
			Nullable: streng.Something("Cherry").Nullable(),
			ExpectedValue:             "Cherry",
			ExpectedOK: true,
		},
		{
			Nullable: streng.Something("dATE").Nullable(),
			ExpectedValue:             "dATE",
			ExpectedOK: true,
		},



		{
			Nullable: streng.Something("Hello world!").Nullable(),
			ExpectedValue:             "Hello world!",
			ExpectedOK: true,
		},



		{
			Nullable: streng.Something("123").Nullable(),
			ExpectedValue:             "123",
			ExpectedOK: true,
		},
		{
			Nullable: streng.Something("123.45").Nullable(),
			ExpectedValue:             "123.45",
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualOK := test.Nullable.Unwrap()

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
