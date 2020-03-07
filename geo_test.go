package geo

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func TestVincenty(t *testing.T) {

	vie := Point{48.110278, 16.569722}
	ewr := Point{40.692480, -74.168686}
	cgn := Point{50.865917, 7.142750}

	cases := []struct {
		x Point
		y Point
		d float64
	}{
		{vie, ewr, 6844607.592055},
		{vie, cgn, 747918.089010},
		{cgn, ewr, 6097517.225179},
	}

	for _, c := range cases {
		d := Vincenty(c.x, c.y)

		if !almostEqual(d, c.d) {
			t.Errorf("wrong distance: got %f, expected %f", d, c.d)
		}
	}
}
