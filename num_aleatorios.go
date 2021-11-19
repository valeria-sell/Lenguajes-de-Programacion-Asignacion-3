package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	//genera semilla en funcion del tiempo
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
	//glc
	r := semilla
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func msg(semilla uint32) func() uint32 {
	//aux del glc
	g := lcg(214013, 2531011, 2048, semilla)
	return func() uint32 {
		return g() % (1 << 16)
	}
}

func gen_array(n int) []int {
	//genera el array de tamano n
	fmt.Println("\n--- Generando: --- ")
	slice := make([]int, n)
	semilla := gen_random_prime()
	//genera semilla basado en el tiempo
	fmt.Println("\n--- Valor semilla: --- ", semilla)
	n_semilla := uint32(semilla)
	msf := msg(n_semilla)
	//genera array de num aleatorios basado en la semilla
	for i := 0; i < n; i++ {
		temp := msf()

		slice[i] = int(temp)
		fmt.Printf("%d ", slice[i])
		//convierte los numeros al rango de 0 a 29 inclusive
		slice[i] = slice[i] % 29

	}
	fmt.Printf("\n")
	return slice
}

