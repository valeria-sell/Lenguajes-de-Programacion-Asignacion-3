package main

import (
	"os"
	"sort"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

// Item2 receive a map to create bar chart->  (label)Key: int, (value)value: float64
func Item2(values map[int]float64, title string) {

	//Sort of the Numbers in the map
	keys := make([]int, 0, len(values))
	for k := range values {
		if values[k] != 0 { // this statement will change on convenience of the chart
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	//Adding the values to the Bar Values List
	var charValues []chart.Value
	for _, k := range keys {
		valV := values[k]
		valL := strconv.Itoa(k)
		charValues = append(charValues, chart.Value{Label: valL, Value: valV})
	}

	graph := chart.BarChart{
		Title: title,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 10, // This value will change on convenience of the chart
		Bars:     charValues,
	}

	f, _ := os.Create(title + "_output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

// To graphic results, receive a map to create bar chart->  (label)Key: string, (value)value: float64
func Item2Palabras(values map[string]float64, title string) {

	//Sort of the Numbers in the map
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	//Adding the values to the Bar Values List
	var charValues []chart.Value
	for l, k := range keys {
		valV := values[k]
		valL := keys[l] //?
		charValues = append(charValues, chart.Value{Label: valL, Value: valV})
	}

	graph := chart.BarChart{
		Title: title,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     charValues,
	}

	f, _ := os.Create(title + "_output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
