package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strconv"
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

func paint(randList []int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bsChart := widgets.NewBarChart()
	bsChart.Data = intToFloat(randList)
	bsChart.Labels = intListString(len(randList))
	bsChart.Title = "Bar Chart QuickSort"
	bsChart.SetRect(5, 5, 100, 25)
	bsChart.BarWidth = 5
	bsChart.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bsChart)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			//return
		}
	}

}

func paint2(slice []int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	qsChart = *widgets.NewBarChart()
	qsChart.Data = intToFloat(slice)
	qsChart.Title = "QuickSort"
	qsChart.SetRect(5, 5, 100, 25)
	qsChart.BarWidth = BAR_WIDTH
	qsChart.BarGap = 0
	qsChart.Labels = intListString(len(slice))
	qsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	qsChart.BarColors = []ui.Color{ui.ColorRed}
	qsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	ui.Render(&qsChart)
	quickSortIterative(slice, 0, len(slice))
}

func update(pair []int, bar widgets.BarChart) {
	swapFloats(&(bar.Data[pair[0]]), &(bar.Data[pair[1]]))
	fmt.Println(bar.Data)
	ui.Render(&bar)
}
