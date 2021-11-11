/*
VERSIÓN RECORTADA POR ITZ PARA OMITIR DETALLES ALGORÍTMICOS
*/

/*
Notas:
1. Para una visualización correcta del gráfico de barras, por favor ejecutar el programa
en una terminal a pantalla completa.
https://www.geeksforgeeks.org/iterative-quick-sort/
https://tecadmin.net/get-current-date-time-golang/
*/

package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"strconv"
	"time"
)



/*
Generates a 3 digit number between 0-599 from the system hour
*/
func generateSeed() int{
/* CÓDIGO OMITIDO */
	return 1 // num
}

/*
creates a N size slice with random numbers based on the linear congruential method
output: slice with N random integers
*/
func randomSlice(size int) []float64{
	var slice = make([]float64, size)
/* CÓDIGO OMITIDO */
	return slice
}

/*
Aux function to swap two numbers
*/
func swap (a *float64, b *float64){
	temp := *a
	*a = *b
	*b = temp
}

/*
Aux function to remove an element from a slice
*/
func remove(slice [][]int, index int) [][]int {
	return append(slice[:index], slice[index+1:]...)
}

//func bubbleSort(slice []float64, pair chan []int){
//	//initTime := time.Now()
//	n := len(slice) - 1
//	for true {
//		/* CÓDIGO OMITIDO */
//		for i := 0; i < n; i++{
//			/* CÓDIGO OMITIDO */
//			if slice[i] > slice[i+1]{
//				pair <- []int{i, i+1}
//				/* CÓDIGO OMITIDO */
//			}
//		}
//		/* CÓDIGO OMITIDO */
//		}
//		/* CÓDIGO OMITIDO */
//	}
//	close(pair)
//	/* CÓDIGO OMITIDO */
//}

//func partition(slice []float64, start int, end int, pair chan []int) int {
//	/* CÓDIGO OMITIDO */
//	for i := start; i < end; i++{
//		/* CÓDIGO OMITIDO */
//		if slice[i] <= pivot{
//			pair <- []int{i,index}
//			/* CÓDIGO OMITIDO */
//		}
//	}
//	pair <- []int{index,end}
//	/* CÓDIGO OMITIDO */
//	return index
//}

/*
Iterative quicksort
 */
//func quickSort(slice []float64, size int, pair chan []int) {
//	/* CÓDIGO OMITIDO */
//	for len(stack) > 0{
//		/* CÓDIGO OMITIDO */
//		pivot := partition(slice, start, end, pair)
//		/* CÓDIGO OMITIDO */
//		if pivot-1 > start {
//			stack = append(stack, []int{start,pivot-1})
//		}
//		/* CÓDIGO OMITIDO */
//		if pivot+1 < end {
//			stack = append(stack, []int{pivot+1,end})
//		}
//		/* CÓDIGO OMITIDO */
//	}
//	close(pair)
//	/* CÓDIGO OMITIDO */
//}




/*
Bubblesort graphic drawer
 */
/*func bsChartDrawer(slice []float64)  {
	bsChart.Data = make([]float64, len(slice))
	copy(bsChart.Data, slice)
	pairsChannel := make(chan []int)
	go bubbleSort(bsChart.Data, pairsChannel)
	for pair := range pairsChannel{
		swap(&bsChart.Data[pair[0]], &bsChart.Data[pair[1]])
		m.Lock()
		ui.Render(&bsChart)
		m.Unlock()
	}
	bsChart.Title = "BubbleSort-Finalizado-" +
		"Tiempo:"+strconv.FormatInt(bsTime.Milliseconds(),10)+"ms-" +
		"Swaps:"+strconv.Itoa(bsSwaps)+"-" +
		"Comparaciones:"+strconv.Itoa(bsComparisons)+"-"+
		"Iteraciones:"+strconv.Itoa(bsIterations)
	m.Lock()
	ui.Render(&bsChart)
	m.Unlock()
}*/

/*
Quicksort graphic drawer
*/
func qsChartDrawer(slice []float64){
	qsChart.Data = make([]float64, len(slice))
	copy(qsChart.Data, slice)
	fmt.Println(qsChart.Data)
	pairsChannel := make(chan []int)
	go quicksort(qsChart.Data, pairsChannel)
	for pair := range pairsChannel{
		fmt.Println(qsChart.Data)
		fmt.Println(qsChart.Data[pair[0]])
		fmt.Println(qsChart.Data[pair[1]])
		time.Sleep(100 * time.Millisecond)
		fmt.Scanln()
		swap(&qsChart.Data[pair[0]], &qsChart.Data[pair[1]])
		m.Lock()
		ui.Render(&qsChart)
		m.Unlock()
	}
	qsChart.Title = "QuickSort-Finalizado-" +
		"Tiempo:"+strconv.FormatInt(qsTime.Milliseconds(),10)+"ms-" +
		"Swaps:"+strconv.Itoa(qsSwaps)+"-" +
		"Comparaciones:"+strconv.Itoa(qsComparisons)+"-"+
		"Iteraciones:"+strconv.Itoa(qsIterations)
	m.Lock()
	ui.Render(&qsChart)
	m.Unlock()
}

func generateLabels(slice []float64) []string {
	var labels = make([]string, len(slice))
	for i := range slice{
		labels[i] = strconv.Itoa(i)
	}
	return labels
}

func initBsChart(slice []float64)  {
	bsChart = *widgets.NewBarChart()
	bsChart.Data = slice
	bsChart.Title = "BubbleSort"
	bsChart.SetRect(0, 0, width, height - 2)
	bsChart.BarWidth = BAR_WIDTH
	bsChart.BarGap = 0
	bsChart.Labels = generateLabels(slice)
	bsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bsChart.BorderBottom = false
	bsChart.BarColors = []ui.Color{ui.ColorRed}
	bsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
}

func initQsChart(slice []float64){
	qsChart = *widgets.NewBarChart()
	qsChart.Data = slice
	qsChart.Title = "QuickSort"
	qsChart.SetRect(0, height-2, width, height*2 - 3)
	qsChart.BarWidth = BAR_WIDTH
	qsChart.BarGap = 0
	qsChart.Labels = generateLabels(slice)
	qsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	qsChart.BarColors = []ui.Color{ui.ColorRed}
	qsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
}