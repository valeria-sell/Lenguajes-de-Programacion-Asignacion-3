package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lxn/win"
	"log"
	"strconv"
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
	width    int = int(win.GetSystemMetrics(win.SM_CXSCREEN) / FONT_WIDTH)
	height   int = int(win.GetSystemMetrics(win.SM_CYSCREEN) / (FONT_HEIGHT * 2))
	bsChart  widgets.BarChart
	qsChart  widgets.BarChart
	ssChart  widgets.BarChart
	hpsChart widgets.BarChart
	isChart  widgets.BarChart

	m sync.Mutex

	bsTime        time.Duration
	bsSwaps       = 0
	bsComparisons = 0
	bsIterations  = 0

	qsTime        time.Duration
	qsSwaps       = 0
	qsComparisons = 0
	qsIterations  = 0
)

func swapFloats(a *float64, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

func intToFloat(intBaseSlice []int) []float64 {
	floatBaseSlice := make([]float64, len(intBaseSlice))
	var v int
	var i int
	for i, v = range intBaseSlice {
		floatBaseSlice[i] = float64(v)
	}
	return floatBaseSlice
}

func intListString(values int) []string {
	intStringSlice := make([]string, values)
	for i := 0; i < values; i++ {
		intStringSlice[i] = strconv.FormatInt(int64(i), 10)
	}
	return intStringSlice
}

func update(pair []int, bar widgets.BarChart) {
	swapFloats(&(bar.Data[pair[0]]), &(bar.Data[pair[1]]))
	//fmt.Println(bar.Data)
	//ui.Render(&bar)
}
func exect(slice []int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
		fmt.Scanln()
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "Lenguajes de Programacion - Asignación # 3 "
	p.Text = "Ana Maria Guevara Rosello\nValeria Sell Saenz"
	p.SetRect(0, 0, 50, 4)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	updateParagraph := func(count int) {
		if count%2 == 0 {
			p.TextStyle.Fg = ui.ColorRed
		} else {
			p.TextStyle.Fg = ui.ColorWhite
		}
	}

	/*
		sparklineData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}

		sl := widgets.NewSparkline()
		sl.Title = "srv 0:"
		sl.Data = sparklineData
		sl.LineColor = ui.ColorCyan
		sl.TitleStyle.Fg = ui.ColorWhite

		sl2 := widgets.NewSparkline()
		sl2.Title = "srv 1:"
		sl2.Data = sparklineData
		sl2.TitleStyle.Fg = ui.ColorWhite
		sl2.LineColor = ui.ColorRed

		slg := widgets.NewSparklineGroup(sl, sl2)
		slg.Title = "Sparkline"
		slg.SetRect(25, 5, 50, 12)*/

	//barchartData := []float64{3, 2, 5, 3, 9 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	labels := intListString(len(slice))
	qsChart = *widgets.NewBarChart()
	qsChart.Title = "QS Chart"
	qsChart.SetRect(0, 4, 202, 13)
	qsChart.Labels = labels
	qsChart.BarColors[0] = ui.ColorGreen
	qsChart.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	qsChart.BarWidth = 1

	ssChart = *widgets.NewBarChart()
	ssChart.Title = "SS Chart"
	ssChart.SetRect(0, 13, 200, 22)
	ssChart.Labels = labels
	ssChart.BarColors[0] = ui.ColorGreen
	ssChart.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	ssChart.BarWidth = 1

	isChart := *widgets.NewBarChart()
	isChart.Title = "IS Chart"
	isChart.SetRect(0, 22, 200, 31)
	isChart.Labels = labels
	isChart.BarColors[0] = ui.ColorGreen
	isChart.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	isChart.BarWidth = 1

	bsChart = *widgets.NewBarChart()
	bsChart.Title = "BS Chart"
	bsChart.SetRect(0, 31, 200, 40)
	bsChart.Labels = labels
	bsChart.BarColors[0] = ui.ColorGreen
	bsChart.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	bsChart.BarWidth = 1

	hpsChart = *widgets.NewBarChart()
	hpsChart.Title = "HPS Chart"
	hpsChart.SetRect(0, 40, 200, 49)
	hpsChart.Labels = labels
	hpsChart.BarColors[0] = ui.ColorGreen
	hpsChart.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	hpsChart.BarWidth = 1

	p2 := widgets.NewParagraph()
	p2.Text = "Hey!\nI am a borderless block!"
	p2.Border = false
	p2.SetRect(50, 10, 75, 10)
	p2.TextStyle.Fg = ui.ColorMagenta

	draw := func(count int) {
		//slg.Sparklines[0].Data = sparklineData[:30+count%50]
		//slg.Sparklines[1].Data = sparklineData[:35+count%50]
		//ui.Render(p, &hpsChart, &bsChart, &isChart, &ssChart, &qsChart,p2)
		ui.Render(p, &ssChart, &qsChart, &isChart, &bsChart, &hpsChart, p2)
		//quickSortIterative(newData, 0, len(slice)-1)
	}

	quickSortData := make([]int, len(slice))
	_ = copy(quickSortData, slice)
	qsChart.Data = intToFloat(quickSortData)

	selectionSortData := make([]int, len(slice))
	_ = copy(selectionSortData, slice)
	ssChart.Data = intToFloat(selectionSortData)

	insertionSortData := make([]int, len(slice))
	_ = copy(insertionSortData, slice)
	isChart.Data = intToFloat(insertionSortData)

	bubbleSortData := make([]int, len(slice))
	_ = copy(bubbleSortData, slice)
	bsChart.Data = intToFloat(bubbleSortData)

	heapSortData := make([]int, len(slice))
	_ = copy(heapSortData, slice)
	hpsChart.Data = intToFloat(heapSortData)

	tickerCount := 1
	draw(tickerCount)
	tickerCount++

	QSChannel := make(chan []int)
	go graphQuickSort(quickSortData, 0, len(quickSortData)-1, QSChannel)
	SSChannel := make(chan []int)
	go graphSelectionSort(selectionSortData, SSChannel)
	ISChannel := make(chan []int)
	go insertionSort(insertionSortData, len(insertionSortData), ISChannel)
	BSChannel := make(chan []int)
	go graphSelectionSort(bubbleSortData, BSChannel)
	HPSChannel := make(chan []int)
	go graphSelectionSort(heapSortData, HPSChannel)

	//uiEvents := ui.PollEvents()
	//ticker := time.NewTicker(time.Second).C
	for {
		time.Sleep(100 * time.Millisecond)
		qsPair, qsOk := <-QSChannel
		ssPair, ssOk := <-SSChannel
		isPair, isOk := <-ISChannel
		bsPair, bsOk := <-BSChannel
		hpsPair, hpsOk := <-HPSChannel

		//fmt.Println(pair)
		tickerCount++
		if qsOk {
			update(qsPair, qsChart)
			//swapFloats(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
			//fmt.Println(swapsQS)
		}
		if ssOk {
			update(ssPair, ssChart)
			//swapFloats(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
			//fmt.Println(swapsQS)
		}
		if isOk {
			update(isPair, isChart)
			//swapFloats(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
			//fmt.Println(swapsQS)
		}
		if bsOk {
			update(bsPair, bsChart)
			//swapFloats(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
			//fmt.Println(swapsQS)
		}
		if hpsOk {
			update(hpsPair, hpsChart)
			//swapFloats(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
			//fmt.Println(swapsQS)
		}

		updateParagraph(tickerCount)
		draw(tickerCount)
	}

	/*for pair := range pairsChannel{
		//fmt.Println(pair)
		swapFloats(&qsChart.Data[pair[0]],&qsChart.Data[pair[1]])
		updateParagraph(tickerCount)
		draw(tickerCount)
		tickerCount++
		fmt.Println(swapsQS)

		if bsChart.Percent <= swapsQS/2 {
			bsChart.Percent = tickerCount
		}
		//fmt.Println(nuevo)
	}*/
	fmt.Scanln()

}
