package main

import (
	"fmt"
	"time"
)

var (
	comparissonsBS = 0
	swapsBS        = 0
	totalTimeBS    time.Duration
)

func graphBubbleSort(randList []int) {
	pairsChannel := make(chan []int)
	go bubbleSort(randList, pairsChannel)
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

func bubbleSort(arr []int, canales chan []int) {
	startTimeBS := time.Now()
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			comparissonsBS++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				canales <- []int{j, j + 1}
				swapsBS++
			}
		}
	}
	close(canales)
	totalTimeBS = time.Since(startTimeBS)
	fmt.Println(totalTimeBS)
}
