package main

import (
	"testing"
)

type testpair struct {
	numset  []float64
	average float64
}

var test = []testpair{
	{[]float64{1, 2, 3}, 2},
	{[]float64{-1, 0, 1}, 0},
	{[]float64{10, 20, 30}, 20},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range test {
		average := Average(pair.numset)
		if average != pair.average {
			t.Error(
				"Pair:", pair.numset,
				"Expected:", pair.average,
				"Got:", average,
			)
		}
	}

}
