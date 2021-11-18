package main

import (
	"time"
)

var (
	comparissonsQS = 0
	swapsQS        = 0
	totalTimeQS    time.Duration
)

//CODE FROM 	https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type Stack []int

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(i int) {
	*s = append(*s, i) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/* This function is same in both iterative and recursive*/
func partition(arr []int, l int, h int, canales chan []int, p *int) {
	x := arr[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		comparissonsQS++
		if arr[j] <= x {
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

	// initialize top of stack
	top := -1

	// push initial values of l and h to stack
	top += 1
	stack.Push(l) //push l
	top += 1
	stack.Push(h) //push h

	// Keep popping from stack while is not empty
	for top >= 0 {
		pairsChannel := make(chan []int)
		// Pop h and l
		top -= 1
		h, _ = stack.Pop() //pop
		top -= 1
		l, _ = stack.Pop() //pop

		// Set pivot element at its correct position
		// in sorted array
		var p int
		go partition(arr, l, h, pairsChannel, &p)
		for pair := range pairsChannel {
			//fmt.Println(arr)
			m.Lock()
			totalTimeQS = time.Since(startTimeQS)
			updater <- pair
			m.Unlock()
		}
		//fmt.Println("p ", p)

		// If there are elements on left side of pivot,
		// then push left side to stack
		if p-1 > l {
			top += 1
			stack.Push(l) //push l
			top += 1
			stack.Push(p - 1) //push p-1
		}

		// If there are elements on right side of pivot,
		// then push right side to stack
		if p+1 < h {
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
