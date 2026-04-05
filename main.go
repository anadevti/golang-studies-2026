package main

import "fmt"

// o arquivo main equivale ao arquivo Program.cs no .Net

func main() {
	//Declarando variáveis
	var a string
	var b int
	var c float64
	var d bool

	fmt.Printf("Meu tipo é %T e meu valor é: '%v'\n", a, a)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", b, b)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", c, c)
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", d, d)

	// Chamando functions
	variables()
	testConst()
	variablesExercise1()
	variablesExercise2()
	variablesExercise3()
	variablesExercise4()
	variablesExercise5()
	variablesExercise6()
	variablesExercise7()
	variablesExercise8()
	exerciseLearning()
	exerciseExerciseTwo()
	exerciseType()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	surprise()
	temperatureClass(float64(0))
	classificationNotes()
	classificationNotesV2()

	// testando frete function:
	frete, err := calcularFrete(10, 200)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Printf("Frete: R$ %.2f\n", frete)
}
