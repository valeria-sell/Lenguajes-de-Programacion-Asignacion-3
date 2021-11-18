package main

import (
	"time"
)

var (
	comparissonsBS = 0
	swapsBS        = 0
	evalsBS        = 0
	totalTimeBS    time.Duration
)

func graphBubbleSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go bubbleSort(randList, pairsChannel)
	for pair := range pairsChannel {
		updater <- pair
	}
	close(updater)
}

func bubbleSort(arr []int, canales chan []int) {
	startTimeBS := time.Now()
	l := len(arr)
	for i := 0; i < l-1; i++ {
		evalsBS++
		for j := 0; j < l-i-1; j++ {
			evalsBS++
			comparissonsBS++
			if arr[j] > arr[j+1] {
				evalsBS++
				arr[j], arr[j+1] = arr[j+1], arr[j]
				totalTimeBS = time.Since(startTimeBS)
				canales <- []int{j, j + 1}
				swapsBS++
			}
		}
	}
	close(canales)
	totalTimeBS = time.Since(startTimeBS)
}
