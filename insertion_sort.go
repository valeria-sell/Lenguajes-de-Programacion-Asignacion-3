package main

import (
	"fmt"
	"time"
)

var (
	comparissonsIS = 0
	swapsIS        = 0
	totalTimeIS    time.Duration
)

func grapInsertioSort(randList []int) {
	pairsChannel := make(chan []int)
	go insertionSort(randList, len(randList), pairsChannel)
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

func insertionSort(arr []int, size int, canales chan []int) {
	startTimeIS := time.Now()
	for i := 1; i < size; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			comparissonsIS++
			arr[j+1] = arr[j]
			canales <- []int{j, j + 1}
			swapsIS++
			j--
		}
		arr[j+1] = key
	}
	close(canales)
	totalTimeIS = time.Since(startTimeIS)
	fmt.Println(totalTimeIS)
}
