package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import (
	"fmt"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lxn/win"
	"sync"
	"time"
)

const (
	BAR_WIDTH       = 3
	FONT_WIDTH      = 8
	FONT_HEIGHT     = 16
	MAX_NUMBER_SIZE = 32
)

var (
	width   int = int(win.GetSystemMetrics(win.SM_CXSCREEN) / FONT_WIDTH)
	height  int = int(win.GetSystemMetrics(win.SM_CYSCREEN) / (FONT_HEIGHT * 2))
	bsChart widgets.BarChart
	qsChart widgets.BarChart
	m       sync.Mutex

	bsTime        time.Duration
	bsSwaps       = 0
	bsComparisons = 0
	bsIterations  = 0

	qsTime        time.Duration
	qsSwaps       = 0
	qsComparisons = 0
	qsIterations  = 0
)

func main() {
	size := 10
	intBaseSlice := gen_array(size)
	fmt.Println(intBaseSlice)
	//heapSort(intBaseSlice, size)
	quickSortIterative(intBaseSlice, 0, size-1)
	fmt.Println(intBaseSlice)

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
