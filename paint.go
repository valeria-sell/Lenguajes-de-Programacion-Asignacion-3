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

func paint(randList []int) widgets.BarChart {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bc := widgets.NewBarChart()
	bc.Data = intToFloat(randList)
	bc.Labels = intListString(len(randList))
	bc.Title = "Bar Chart QuickSort"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			//return
		}
	}

	return *bc
}

func update(pair []int, bar widgets.BarChart) {
	swapFloats(&(bar.Data[pair[0]]), &(bar.Data[pair[1]]))
	fmt.Println(bar.Data)
	ui.Render(&bar)
}
