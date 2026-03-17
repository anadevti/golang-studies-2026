package main

import (
	"fmt"
)

// o arquivo main equivale ao arquivo Program.cs no .Net

func main() {
	// Declarando variáveis
	var a string
	var b int
	var c float64
	var d bool

	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", a, a)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", b, b)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", c, c)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", d, d)

	// Chamando functions
	testConst()
	variablesExercise1()
	variablesExercise2()
	variablesExercise3()
	variablesExercise4()
	variablesExercise5()
}
