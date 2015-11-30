package maths_test

import (
	"testing"

	"github.com/sasimpson/gotilites"
)

func TestRoundDown(t *testing.T) {
	cases := []struct {
		in, places, result int
	}{
		{12000, 2, 12000},
		{12200, 2, 12200},
		{12249, 2, 12200},
	}
	for _, c := range cases {
		got := maths.RoundInt(c.in, c.places)
		if got != c.result {
			t.Errorf("didn't round down correctly %d to %d places should be %d", c.in, c.places, c.result)
		}
	}

}

func TestRoundUp(t *testing.T) {
	cases := []struct {
		in, places, result int
	}{
		{12150, 2, 12200},
		{12500, 3, 13000},
		{13781, 4, 10000},
	}
	for _, c := range cases {
		got := maths.RoundInt(c.in, c.places)
		if got != c.result {
			t.Errorf("didn't round up correctly %d to %d places should be %d", c.in, c.places, c.result)
		}
	}
}
