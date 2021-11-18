package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

var (
	//Indicadores de las coordenadas Y para las funciones graficadoras
	prvYSize  = 0
	nextYSize = 6

	//Párrafos correspondientes a la información del programa
	p      widgets.Paragraph // Párrafo de título
	pTable widgets.Paragraph // Indicadores
	pqs    widgets.Paragraph // Info del QuickSort
	pss    widgets.Paragraph // Info del SelectionSort
	pis    widgets.Paragraph // Info del BubbleSort
	pbs    widgets.Paragraph // Info del InsertionSort
	phs    widgets.Paragraph // Info del HeapSort

	// Graficos de barras correspondientes a cada sort
	bsChart  widgets.BarChart // BubbleSort
	qsChart  widgets.BarChart // QuickSort
	ssChart  widgets.BarChart // SelectionSort
	hpsChart widgets.BarChart //HeapSort
	isChart  widgets.BarChart //InsertionSort

	m sync.Mutex //encargado de sincronizar los eventos y canales
)

// swapFloats: realiza un intercambio entre 2 elementos flotantes
func swapFloats(a *float64, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

// intToFloat: realiza un intercambio entre 2 elementos enteros
func intToFloat(intBaseSlice []int) []float64 {
	floatBaseSlice := make([]float64, len(intBaseSlice))
	var v int
	var i int
	for i, v = range intBaseSlice {
		floatBaseSlice[i] = float64(v)
	}
	return floatBaseSlice
}

// intListString: genera una lista de strings, con la cantidad de elementos dada
func intListString(values int) []string {
	intStringSlice := make([]string, values)
	for i := 0; i < values; i++ {
		intStringSlice[i] = strconv.FormatInt(int64(i), 10)
	}
	return intStringSlice
}

// update: función encargada de realizar los intercambios de elementos dentro del gráfico correspondiente
func update(pair []int, bar *widgets.BarChart) {
	swapFloats(&(bar.Data[pair[0]]), &(bar.Data[pair[1]]))
}

// paragraphs: función encargada de inicializar los valores de los párrafos a graficar, así como su diseño inicial
func paragraphs() {
	prvXSize := 0
	nextXSize := 45
	distance := 15

	p = *widgets.NewParagraph()                              // Se crea el nuevo párrafo
	p.Title = "Lenguajes de Programacion - Asignación # 3 "  //Asignación del título
	p.Text = "Ana Maria Guevara Rosello\nValeria Sell Saenz" //Asignación del texto
	p.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)      //Ubicación según pares ordenados (x1,y1,x2,y2)
	p.TextStyle.Fg = ui.ColorWhite                           //Color del texto
	p.BorderStyle.Fg = ui.ColorCyan                          //Color del borde

	prvXSize = nextXSize            // Se actualiza el ultimo valor de X2, como X1
	nextXSize = prvXSize + distance // Se modifica el X2 según el tamaño deseado
	pTable = *widgets.NewParagraph()
	pTable.Title = "Sorts"
	pTable.Text = "Comparaciones\nIntercambios\nOpElementales\nTiempo"
	pTable.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	pTable.TextStyle.Fg = ui.ColorWhite
	pTable.BorderStyle.Fg = ui.ColorCyan

	prvXSize = nextXSize
	nextXSize = prvXSize + distance
	pqs = *widgets.NewParagraph()
	pqs.Title = "QuickSort" //color blanco al inicio, y parpadear en verde finalizar
	pqs.Text = strconv.Itoa(comparissonsQS) + "\n" + strconv.Itoa(swapsQS) + "\n" + strconv.Itoa(evalsQS) + "\n" + totalTimeQS.String()
	pqs.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	pqs.TitleStyle.Fg = ui.ColorWhite
	pqs.TextStyle.Fg = ui.ColorWhite
	pqs.BorderStyle.Fg = ui.ColorCyan

	prvXSize = nextXSize
	nextXSize = prvXSize + distance
	pss = *widgets.NewParagraph()
	pss.Title = "Selection" //color blanco al inicio, y parpadear en verde finalizar
	pss.Text = strconv.Itoa(comparissonsSS) + "\n" + strconv.Itoa(swapsSS) + "\n" + strconv.Itoa(evalsSS) + "\n" + totalTimeSS.String()
	pss.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	pss.TitleStyle.Fg = ui.ColorWhite
	pss.TextStyle.Fg = ui.ColorWhite
	pss.BorderStyle.Fg = ui.ColorCyan

	prvXSize = nextXSize
	nextXSize = prvXSize + distance
	pis = *widgets.NewParagraph()
	pis.Title = "Insertion" //color blanco al inicio, y parpadear en verde finalizar
	pis.Text = strconv.Itoa(comparissonsIS) + "\n" + strconv.Itoa(swapsIS) + "\n" + strconv.Itoa(evalsIS) + "\n" + totalTimeIS.String()
	pis.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	pis.TitleStyle.Fg = ui.ColorWhite
	pis.TextStyle.Fg = ui.ColorWhite
	pis.BorderStyle.Fg = ui.ColorCyan

	prvXSize = nextXSize
	nextXSize = prvXSize + distance
	pbs = *widgets.NewParagraph()
	pbs.Title = "Bubble" //color blanco al inicio, y parpadear en verde finalizar
	pbs.Text = strconv.Itoa(comparissonsBS) + "\n" + strconv.Itoa(swapsBS) + "\n" + strconv.Itoa(evalsBS) + "\n" + totalTimeBS.String()
	pbs.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	pbs.TitleStyle.Fg = ui.ColorWhite
	pbs.TextStyle.Fg = ui.ColorWhite
	pbs.BorderStyle.Fg = ui.ColorCyan

	prvXSize = nextXSize
	nextXSize = prvXSize + distance
	phs = *widgets.NewParagraph()
	phs.Title = "Heap" //color blanco al inicio, y parpadear en verde finalizar
	phs.Text = strconv.Itoa(comparissonsHS) + "\n" + strconv.Itoa(swapsHS) + "\n" + strconv.Itoa(evalsHPS) + "\n" + totalTimeHS.String()
	phs.SetRect(prvXSize, prvYSize, nextXSize, nextYSize)
	phs.TitleStyle.Fg = ui.ColorWhite
	phs.TextStyle.Fg = ui.ColorWhite
	phs.BorderStyle.Fg = ui.ColorCyan

	prvYSize = nextYSize
}

// barchars: función encargada de inicializar los valores de los gráficos de barras, así como su diseño inicial
func barchars(size int) {
	labels := intListString(size)

	//Se utiliza un tamaño dinámico para el ancho de los gráficos de barras
	fullYSize := 202
	ySize := int(math.Ceil((float64(size)/100.0)*float64(fullYSize))) + 1

	qsChart = *widgets.NewBarChart() // Se crea una nueva BarChart
	qsChart.Title = "QS Chart"       // Se define su título
	nextYSize = prvYSize + 9
	qsChart.SetRect(0, prvYSize, ySize, nextYSize)                                          //Ubicación según pares ordenados (x1,y1,x2,y2)
	qsChart.Labels = labels                                                                 // Se usan los labels creados previamente
	qsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite), ui.NewStyle(ui.ColorCyan)} //Se define el estilo y color de los labels
	qsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}                              // Se define el estilo y color de los números en las barras
	qsChart.BarWidth = 1                                                                    // Ancho de las barras

	prvYSize = nextYSize
	nextYSize = prvYSize + 9
	ssChart = *widgets.NewBarChart()
	ssChart.Title = "SS Chart"
	ssChart.SetRect(0, prvYSize, ySize, nextYSize)
	ssChart.Labels = labels
	ssChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite), ui.NewStyle(ui.ColorCyan)}
	ssChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	ssChart.BarWidth = 1

	prvYSize = nextYSize
	nextYSize = prvYSize + 9
	isChart = *widgets.NewBarChart()
	isChart.Title = "IS Chart"
	isChart.SetRect(0, prvYSize, ySize, nextYSize)
	isChart.Labels = labels
	isChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite), ui.NewStyle(ui.ColorCyan)}
	isChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	isChart.BarWidth = 1

	prvYSize = nextYSize
	nextYSize = prvYSize + 9
	bsChart = *widgets.NewBarChart()
	bsChart.Title = "BS Chart"
	bsChart.SetRect(0, prvYSize, ySize, nextYSize)
	bsChart.Labels = labels
	bsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite), ui.NewStyle(ui.ColorCyan)}
	bsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	bsChart.BarWidth = 1

	prvYSize = nextYSize
	nextYSize = prvYSize + 9
	hpsChart = *widgets.NewBarChart()
	hpsChart.Title = "HPS Chart"
	hpsChart.SetRect(0, prvYSize, ySize, nextYSize)
	hpsChart.Labels = labels
	hpsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite), ui.NewStyle(ui.ColorCyan)}
	hpsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	hpsChart.BarWidth = 1
}

// graficar: Función encargada de llevar a cabo todos los procesos de unión entre los gráficos y las gorrutinas de ordenamientos
// Esta función recibe unicamente la lista de elementos original para relizar los ordenamientos
func graficar(slice []int) {
	// En caso de que haya un error en la graficación, el programa lo indicará y esperará un input
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
		fmt.Scanln()
	}
	defer ui.Close()

	//función utilizada para modificar el color de un título en caso de éxito en su ordenamiento
	successTitleColor := func(count int, p *widgets.Paragraph) {
		if count%2 == 0 {
			p.TitleStyle.Fg = ui.ColorBlack
		} else {
			p.TitleStyle.Fg = ui.ColorGreen
		}
	}

	// función encargada de actualizar los valores de los párrafos, estos corresponden a las estadísticas de los ordenamientos
	updateParagraph := func(count int) {
		pqs.Text = strconv.Itoa(comparissonsQS) + "\n" + strconv.Itoa(swapsQS) + "\n" + strconv.Itoa(evalsQS) + "\n" + totalTimeQS.String()
		pss.Text = strconv.Itoa(comparissonsSS) + "\n" + strconv.Itoa(swapsSS) + "\n" + strconv.Itoa(evalsSS) + "\n" + totalTimeSS.String()
		pis.Text = strconv.Itoa(comparissonsIS) + "\n" + strconv.Itoa(swapsIS) + "\n" + strconv.Itoa(evalsIS) + "\n" + totalTimeIS.String()
		pbs.Text = strconv.Itoa(comparissonsBS) + "\n" + strconv.Itoa(swapsBS) + "\n" + strconv.Itoa(evalsBS) + "\n" + totalTimeBS.String()
		phs.Text = strconv.Itoa(comparissonsHS) + "\n" + strconv.Itoa(swapsHS) + "\n" + strconv.Itoa(evalsHPS) + "\n" + totalTimeHS.String()
	}

	//Se lleva a cabo la inicialización de los párrafos y los gráficos de barras
	paragraphs()
	barchars(len(slice))

	// función encargada de renderizar, es decir, dibujar los gráficos en consola
	draw := func() {
		ui.Render(&p, &pTable, &pqs, &pss, &pis, &pbs, &phs, &qsChart, &ssChart, &isChart, &bsChart, &hpsChart)
	}

	//Se realiza una copia de la lista original de números para cada ordenamiento
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

	//Se inicia un contador que nos será útil para la actualización de datos
	tickerCount := 1
	draw()

	// Se inician los canales, junto con sus recibidores, que podrán recibir datos de tipo []int
	QSChannel := make(chan []int)
	go graphQuickSort(quickSortData, 0, len(quickSortData)-1, QSChannel)

	SSChannel := make(chan []int)
	go graphSelectionSort(selectionSortData, SSChannel)

	ISChannel := make(chan []int)
	go insertionSort(insertionSortData, len(insertionSortData), ISChannel)

	BSChannel := make(chan []int)
	go graphBubbleSort(bubbleSortData, BSChannel)

	HPSChannel := make(chan []int)
	go graphHeapSort(heapSortData, HPSChannel)

	// Se hace uso de un ciclo infinito para la visualización de las gráficas
	for {
		// Este for tendrá una espera de 0.3 segundos entre cada interecambio de valores en los ordenamientos
		time.Sleep(300 * time.Millisecond)

		//Cada canal recibe a su propio conjunto de datos, a travez de <-'SChannel , lo enviado por su algoritmo
		// el primer ('sPair) valor corresponde a un par de números, que serán las ubicaciones de los números a intercambiar
		// el segundo valor ('sOkay) será un boleano, encargado de decirnos si el canal se encuentra abierto o cerrado
		qsPair, qsOk := <-QSChannel
		ssPair, ssOk := <-SSChannel
		isPair, isOk := <-ISChannel
		bsPair, bsOk := <-BSChannel
		hpsPair, hpsOk := <-HPSChannel

		tickerCount++
		// Si el canal se encuentre abierto, quiere decir que hay un par de elementos a intercambiar
		if qsOk {
			// Se realiza el intercambio
			update(qsPair, &qsChart)
			//En caso de que el canal se encuentre cerrado, quiere decir que el ordenamiento ha terminado
		} else {
			// Se usa un parpadeo verde en el título del ordenamiento para indicar que este fue éxitoso y ha terminado
			successTitleColor(tickerCount, &pqs)
		}
		if ssOk {
			update(ssPair, &ssChart)
		} else {
			successTitleColor(tickerCount, &pss)
		}
		if isOk {
			update(isPair, &isChart)
		} else {
			successTitleColor(tickerCount, &pis)
		}
		if bsOk {
			update(bsPair, &bsChart)
		} else {
			successTitleColor(tickerCount, &pbs)
		}
		if hpsOk {
			update(hpsPair, &hpsChart)
		} else {
			successTitleColor(tickerCount, &phs)
		}
		// Al terminar de evaluar todos los canales, se actualizan las estadísticas y empieza de nuevo el ciclo
		updateParagraph(tickerCount)
		draw()
	}
	fmt.Scanln()

}
