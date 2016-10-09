package main

import "math/rand"

func computeEmptyIndeces(emptyPercent float32, count int) (map[int]bool, error) {
	// precondition: <0,100> checked by json schema already

	indeces := make(map[int]bool, count)

	if emptyPercent != 0 {
		if emptyPercent == 100 {
			for i := 0; i < count; i++ {
				indeces[i] = true
			}
		} else {
			randInts := rand.Perm(count)

			emptyCount := float32(count) * emptyPercent / 100
			for _, val := range randInts[:int(emptyCount)] {
				indeces[val] = true
			}
		}
	}
	return indeces, nil
}
