package main

import (
	"time"
)

var (
	comparissonsQS = 0
	swapsQS        = 0
	evalsQS        = 0
	totalTimeQS    time.Duration
)

//CODE FROM 	https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type Stack []int

// IsEmpty revisa si el stack esta vacio
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push introduce un nuevo valor al stack
func (s *Stack) Push(i int) {
	*s = append(*s, i) // agrega el nuevo elemento al final del stack
}

// Pop Hace pop al primer elemento de la pila. Retorna false si esta vacia
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // obtiene indice del primer elemento
		element := (*s)[index] // indexa al slice y obtiene el elemento.
		*s = (*s)[:index]      // retira elemento del stack
		return element, true
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
//funcion de particion
func partition(arr []int, l int, h int, canales chan []int, p *int) {
	x := arr[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		comparissonsQS++
		evalsQS++
		if arr[j] <= x {
			evalsQS++
			i++
			swap(&arr[i], &arr[j])
			//fmt.Println("i ", i, " j ", j)
			canales <- []int{i, j}
			swapsQS++
		}
	}
	swap(&arr[i+1], &arr[h])
	swapsQS++

	canales <- []int{i + 1, h}
	*p = i + 1
	//fmt.Println("p go", *p)
	close(canales)
	//return i + 1
}

func graphQuickSort(arr []int, l int, h int, updater chan []int) {
	// Create an auxiliary stack
	//len := h - l + 1
	startTimeQS := time.Now()

	original := make([]int, len(arr))
	copy(original, arr)
	var stack Stack

	//paint2(original)

	// inicializa stack
	top := -1

	// push a l y j iniciales
	top += 1
	stack.Push(l) //push l
	top += 1
	stack.Push(h) //push h

	// condicion para seguir haciendo pop a la pila hasta que este vacia
	for top >= 0 {
		evalsQS++
		pairsChannel := make(chan []int)
		// Pop h y l
		top -= 1
		h, _ = stack.Pop() //pop
		top -= 1
		l, _ = stack.Pop() //pop

		// pone el pivote en la condicion correcta
		var p int
		go partition(arr, l, h, pairsChannel, &p)
		for pair := range pairsChannel {
			//envia datos al canal
			//fmt.Println(arr)
			m.Lock()
			totalTimeQS = time.Since(startTimeQS)
			updater <- pair
			m.Unlock()
		}
		//fmt.Println("p ", p)

		// si hay elementos a la izq del pivote,
		// hace push al lado izq al stack
		if p-1 > l {
			evalsQS++
			top += 1
			stack.Push(l) //push l
			top += 1
			stack.Push(p - 1) //push p-1
		}

		// si hay elementos a la derecho del pivote,
		// hace push al lado derecho al stack
		if p+1 < h {
			evalsQS++
			top += 1
			stack.Push(p + 1) //push p+1
			top += 1
			stack.Push(h) //push h
		}
	}
	totalTimeQS = time.Since(startTimeQS)
	//fmt.Println(totalTimeQS)
	close(updater)
}
