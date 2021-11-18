package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

func main() {
	size := 10
	intBaseSlice := gen_array(size)
	nuevo := make([]int, size)
	_ = copy(nuevo, intBaseSlice)
	graficar(nuevo)
}
