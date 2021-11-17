package main

import (
	"fmt"
	"time"
)

var (
	comparissonsSS = 0
	swapsSS        = 0
	totalTimeSS    time.Duration
)

func graphSelectionSort(randList []int) {
	pairsChannel := make(chan []int)
	go selectionsort(randList, pairsChannel)
	for pair := range pairsChannel {
		fmt.Println(pair)
		/*m.Lock()
		//update(pair,*bc)

		swapFloats(&(bsChart.Data[pair[0]]), &(bsChart.Data[pair[1]]))
		fmt.Println(bsChart.Data)
		ui.Render(&bsChart)
		time.Sleep(100 * time.Millisecond)
		m.Unlock()*/
	}
	fmt.Println(randList)

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
		canales <- []int{i, minIdx}
		swapsSS++
	}
	close(canales)
	totalTimeSS = time.Since(startTimeSS)
	fmt.Println(totalTimeSS)
}
