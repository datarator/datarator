package main

import "testing"

func TestComputeEmptyIndeces(t *testing.T) {
	var tests = []struct {
		emptyPercent    float32
		count           int
		emptyIndecesCnt int
	}{
		{
			emptyPercent:    0,
			count:           10,
			emptyIndecesCnt: 0,
		},
		{
			emptyPercent:    100,
			count:           10,
			emptyIndecesCnt: 10,
		},
		{
			emptyPercent:    50,
			count:           10,
			emptyIndecesCnt: 5,
		},
		{
			emptyPercent:    45,
			count:           10,
			emptyIndecesCnt: 4,
		},
		{
			emptyPercent:    33,
			count:           10,
			emptyIndecesCnt: 3,
		},
	}

	for _, test := range tests {
		actual, _ := computeEmptyIndeces(test.emptyPercent, test.count)
		actualEmptyIndecesCnt := 0
		for i, _ := range actual {
			if actual[i] {
				actualEmptyIndecesCnt++
			}
		}
		if actualEmptyIndecesCnt != test.emptyIndecesCnt {
			t.Fatalf("Expected: %v\nActual: %v", test.emptyIndecesCnt, actualEmptyIndecesCnt)
		}
	}
}
