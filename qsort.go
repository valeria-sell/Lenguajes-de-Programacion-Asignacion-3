package main

import (
	"fmt"
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
func partition(arr []int, l int, h int, canales chan []int) int {
	x := arr[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		if arr[j] <= x {
			i++
			swap(&arr[i], &arr[j])
			canales <- []int{j, i + 1}
		}
	}
	swap(&arr[i+1], &arr[h])
	canales <- []int{i + 1, h}
	fmt.Println(&canales)
	return i + 1
}

/* A[] --> Array to be sorted,
   l --> Starting index,
   h --> Ending index */

func quickSortIterative(arr []int, l int, h int) {
	// Create an auxiliary stack
	//len := h - l + 1
	var stack Stack

	// initialize top of stack
	top := -1

	// push initial values of l and h to stack
	top += 1
	stack.Push(l) //push l
	top += 1
	stack.Push(h) //push h
	pairsChannel := make(chan []int)

	// Keep popping from stack while is not empty
	for top >= 0 {
		// Pop h and l
		top -= 1
		h, _ = stack.Pop() //pop
		top -= 1
		l, _ = stack.Pop() //pop

		// Set pivot element at its correct position
		// in sorted array
		p := partition(arr, l, h, pairsChannel)

		for pair := range pairsChannel {
			fmt.Println(pair)
		}

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
}

/*
// Driver code
func main() int {

int n = 5;
int arr[n] = { 4, 2, 6, 9, 2 };

quickSort(arr, 0, n - 1);

for (int i = 0; i < n; i++) {
out << arr[i] << " ";
}

return 0;
}*/
