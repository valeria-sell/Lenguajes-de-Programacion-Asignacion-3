package main

import (
	"fmt"
	"math/rand"
	"time"
)

// prime returns true if n is a prime number.
func is_prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gen_random_prime() int {
	min := 11
	max := 101
	rand.Seed(time.Now().UnixNano())
	randint := rand.Intn(max-min) + min
	val := is_prime(4)
	for !val { // emulates while
		// do something
		randint = rand.Intn(max-min) + min
		val = is_prime(randint)
	}
	return randint
}

func lcg(a, c, m, semilla uint32) func() uint32 {
	r := semilla
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func msg(semilla uint32) func() uint32 {
	g := lcg(214013, 2531011, 2048, semilla)
	return func() uint32 {
		return g() % (1 << 16)
	}
}

func gen_array(n int) []int {
	fmt.Println("\n--- Generando: --- ")
	slice := make([]int, n)
	semilla := gen_random_prime()
	fmt.Println("\n--- Valor semilla: --- ", semilla)
	n_semilla := uint32(semilla)
	msf := msg(n_semilla)
	for i := 0; i < n; i++ {
		temp := msf()

		slice[i] = int(temp)
		fmt.Printf("%d ", slice[i])
		slice[i] = slice[i] % 29

	}
	fmt.Printf("\n")
	return slice
}

/*func main() {
	A := gen_array(100)
	B := gen_array(400)
	C := gen_array(600)
	D := gen_array(800)
	E := gen_array(1000)

	for _, value := range A {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range B {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range C {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range D {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range E {
		fmt.Printf("%d ", value)
	}
}*/
