package main

import (
	"time"
)

var (
	comparissonsHS = 0
	swapsHS        = 0
	evalsHPS       = 0
	totalTimeHS    time.Duration
)

func graphHeapSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go heapSort(randList, len(randList), pairsChannel)
	for pair := range pairsChannel {
		updater <- pair
	}
	close(updater)
}

func buildMaxHeap(arr []int, n int, canales chan []int) {
	for i := 1; i < n; i++ {
		//si el nodo hijo es mayor al nodo padre
		comparissonsHS++
		evalsHPS++
		if arr[i] > arr[(i-1)/2] {
			evalsHPS++
			j := i
			// intercambia nodo padre y nodo hijo hasta que el padre sea menor
			for arr[j] > arr[(j-1)/2] {
				evalsHPS++
				swap(&arr[j], &arr[(j-1)/2])
				swapsHS++
				canales <- []int{j, (j - 1) / 2}
				j = (j - 1) / 2
			}
		}
	}
}

func heapSort(arr []int, n int, canales chan []int) {
	startTimeHS := time.Now()
	//heap auxiliar
	buildMaxHeap(arr, n, canales)
	//fmt.Println(arr)
	for i := n - 1; i > 0; i-- {
		// intercambia los valores del primer elemento indexado
		// con el ultimo
		swap(&arr[0], &arr[i])
		totalTimeHS = time.Since(startTimeHS)
		canales <- []int{0, i}
		swapsHS++
		evalsHPS++
		// mantiene la integridad el heap durante cada intercambio
		j := 0
		index := 0

		for {
			//adaptacion del while loop para go
			evalsHPS++
			index = 2*j + 1
			if index > i {
				//condicion de parada del ciclo
				evalsHPS++
				break
			}
			// si el hijo izq es menor que el hijo derecho
			// apunta variable del indice al hijo derecho
			if arr[index] < arr[index+1] && index < (i-1) {
				evalsHPS++
				index++
			}
			// si el nodo padre es menor al nodo hijo
			// intercambia nodo padre con el nodo hijo de mayor valor
			comparissonsHS++
			if arr[j] < arr[index] && index < i {
				evalsHPS++
				swap(&arr[j], &arr[index])
				swapsHS++
				totalTimeHS = time.Since(startTimeHS)
				canales <- []int{j, index}
			}
			j = index
		}
	}
	close(canales)
	totalTimeHS = time.Since(startTimeHS)
	//fmt.Println(totalTimeHS)
}
