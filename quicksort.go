package main

import (
	"math/rand"
)

//Own Code
func quicksort(a []float64, b chan []int) chan []int {
	if len(a) < 2 {
		b <- []int{0, len(a)}
		return b
	}
	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}

	}
	a[left], a[right] = a[right], a[left]

	go quicksort(a[:left], b)
	go quicksort(a[left+1:], b)
	b <- []int{0, len(a)}
	return b
}
