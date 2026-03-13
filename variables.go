package main

import (
	"fmt"
)

func variables() {
	// Declarando variáveis
	var a string
	var b int
	var c float64
	var d bool

	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", a, a)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", b, b)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", c, c)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", d, d)

	// Declarando uma lista de variaveis
	var we, are, gophers string

	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", we, we)
	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", are, are)
	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", gophers, gophers)
}
