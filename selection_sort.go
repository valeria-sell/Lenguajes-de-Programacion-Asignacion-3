package main

import (
	"time"
)

var (
	comparissonsSS = 0
	swapsSS        = 0
	totalTimeSS    time.Duration
)

func graphSelectionSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go selectionsort(randList, pairsChannel)
	for pair := range pairsChannel {
		updater <- pair

	}
	close(updater)

}

func selectionsort(items []int, canales chan []int) {
	startTimeSS := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			comparissonsSS++
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
		totalTimeSS = time.Since(startTimeSS)
		canales <- []int{i, minIdx}
		swapsSS++
	}
	close(canales)
	totalTimeSS = time.Since(startTimeSS)
}
