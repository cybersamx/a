package a

import (
	"math"
	"testing"
)

var tcs = []struct {
	m            []float64
	stddevExpect float64
}{
	{[]float64{1, 2, 3, 4}, 1.11803399},
	{[]float64{610, 450, 160, 420, 310}, 149.799866},
}

func compare(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}

func TestStdDev(t *testing.T) {
	for _, tc := range tcs {
		res := StdDev(tc.m)
		expect := tc.stddevExpect

		if !compare(res, expect) {
			t.Errorf("StdDev: %f != %f", res, expect)
		}
	}
}
