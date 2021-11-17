package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import (
	"fmt"
)

func main() {
	size := 100
	intBaseSlice := gen_array(size)
	nuevo := make([]int, size)
	_ = copy(nuevo, intBaseSlice)
	//fmt.Println(nuevo, n)
	//heapSort(intBaseSlice, size)
	//quickSortIterative(intBaseSlice, 0, size-1)
	//grapInsertioSort(intBaseSlice)
	//exec(intBaseSlice)
	/*pairsChannel := make(chan []int)
	go quickSortIterative(intBaseSlice, 0, size-1, pairsChannel)
	fmt.Println(intBaseSlice)
	for pair := range pairsChannel{
		fmt.Println(pair)
		swap(&nuevo[pair[0]],&nuevo[pair[1]])
		fmt.Println(nuevo)
	}*/

	exect(nuevo)
	//fmt.Println(intBaseSlice)

	/*	floatBaseSlice := make([]float64, len(intBaseSlice))
		pairsChannel := make(chan []int)
		var v int
		var i int
		for i, v = range intBaseSlice {
			floatBaseSlice[i] = float64(v)
		}
		fmt.Println(floatBaseSlice)
		go quicksort(floatBaseSlice, pairsChannel)
		cont := 0
		for pair := range pairsChannel{
			fmt.Println(floatBaseSlice)
			fmt.Println(pair)
			cont++
		}*/

	/*Block{
		Try: func() {
			barNumber := width / BAR_WIDTH - 1
			fmt.Print("Indique la cantidad de numeros(Se recomienda " + strconv.Itoa(barNumber) +" maximo para una visualizacion correcta): ")
			var size int
			fmt.Scanln(&size)
			if err := ui.Init(); err != nil {
				log.Fatalf("failed to initialize termui: %v", err)
			}
			defer ui.Close()


			fmt.Scanln()
			intBaseSlice := gen_array(size)
			floatBaseSlice := make([]float64, len(intBaseSlice))
			var v int
			var i int
			for i, v = range intBaseSlice {
				floatBaseSlice[i] = float64(v)
			}
			fmt.Scanln()
			//initBsChart(floatBaseSlice)
			initQsChart(floatBaseSlice)
			//ui.Render(&bsChart)
			ui.Render(&qsChart)
			//go bsChartDrawer(floatBaseSlice)
			qsChartDrawer(floatBaseSlice)
			fmt.Scanln() //end until any key is pressed
			uiEvents := ui.PollEvents()
			for {
				e := <-uiEvents
				switch e.ID {
				case "q", "<C-c>":
					return
				}
			}

			ui.Close()
			Throw("Oh,...sh...")
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()*/
	fmt.Println("We did it!")

}
