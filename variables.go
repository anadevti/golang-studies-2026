package main

import (
	"fmt"
)

// Exercicios feitos referente a: https://womenwhogocwb.gitbook.io/letsgo/tipos-de-dados-basicos/exercicios

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

func variablesExercise1() {
	var test int
	test = 10
	fmt.Printf("test = %d\n", test)
}

func variablesExercise2() {
	var test int
	test = 20
	fmt.Printf("test = %d\n", test)
}

func variablesExercise3() {
	var test int
	test = 30
	fmt.Printf("test = %d\n", test)
}

func variablesExercise4() {
	a := 230
	b := 27
	sum := a + b
	minn := a - b
	fmt.Printf("a = %d, b = %d, sum = %d\n", a, b, sum)
	fmt.Printf("a = %d, b = %d, sum = %d\n", a, b, minn)
}

func variablesExercise5() {
	var precoDaBanana, precoDaCerveja, precoDoAbacate, precoDoSalgadinho float64 // lista

	precoDaBanana = 1.25 // atribuindo valores aos itens da lista
	precoDaCerveja = 2.98
	precoDoAbacate = 4.59
	precoDoSalgadinho = 7.29

	var pesoBanana, quantidadeCerveja, pesoAbacate, quantidadeSalgadinho float64

	pesoBanana = 2.170
	quantidadeCerveja = 6
	pesoAbacate = 5.650
	quantidadeSalgadinho = 3

	// O * aqui está fazendo multiplicação (operador aritmético), não é ponteiro.
	// O ponteiro em Go é declarado junto a um tipo ou variável, por exemplo: *int, &variavel.
	// Quando aparece entre dois valores numéricos, é sempre multiplicação.

	valorDaCompra := (precoDaBanana * pesoBanana) + (precoDaCerveja * quantidadeCerveja) + // calculos
		+(precoDoAbacate * pesoAbacate) + (quantidadeSalgadinho * precoDoSalgadinho)

	fmt.Printf("O valor da compra deu R$ %v reais. O Salgadinho custou: %v a unidade", valorDaCompra, precoDoSalgadinho)
}

func variablesExercise6() {
	var test string
	test = "Meu\nNome é Ana e\n"

	testString := `minha cor favorita é azul`
	fmt.Printf("test = %s\n", test+testString)
}

func variablesExercise7() {
	a := 10
	b := 20
	c := 30
	d := 40
	e := 50
	fmt.Println(a == b)
	fmt.Println(b != a)
	fmt.Println(c > d)
	fmt.Println(a < e)
	fmt.Println(d <= 50)
}

func variablesExercise8() {
	balance := 25
	haveMovements := true

	if balance >= 30 && haveMovements {
		fmt.Println("Have movements")
	} else {
		fmt.Println("No Have movements")
	}
}
