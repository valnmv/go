package temperature

import (
	"testing"
)

type temperatureTest struct {
	i        Temperature
	expected Temperature
}

var CtoFTests = []temperatureTest{
	{4.1, 39.38},
	{10, 50},
	{-10, 14},
}

func testCtoF(t *testing.T) {
	for _, tt := range CtoFTests {
		actual := CtoF(tt.i)
		if actual != tt.expected {
			t.Errorf("expected %v, actual %v", tt.expected, actual)
		}
	}
}
