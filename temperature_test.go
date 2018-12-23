package temperature

import (
	"testing"
)

type temperatureTest struct {
	i    Temperature
	want Temperature
}

var CtoFTests = []temperatureTest{
	{4.1, 39.38},
	{10, 50},
	{-10, 14},
}

func TestCtoF(t *testing.T) {
	for _, tt := range CtoFTests {
		got := CtoF(tt.i)
		if got != tt.want {
			t.Errorf("expected %v, got %v", tt.want, got)
		}
	}
}

var FtoCTests = []temperatureTest{
	{32, 0},
	{86, 30},
	{140, 60},
}

func TestFtoC(t *testing.T) {
	for _, tt := range FtoCTests {
		got := FtoC(tt.i)
		if got != tt.want {
			t.Errorf("expected %v, got %v", tt.want, got)
		}
	}
}
