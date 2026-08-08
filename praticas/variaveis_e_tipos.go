package praticas

import (
	"fmt"
)

/*
Nível 1 — Variáveis e tipos
Exercício 1 — Cadastro simples

Crie um programa que declare:

nome como string
idade como int
altura como float64
estudando como bool

Depois, imprima todos os valores.

Pratique:

var
inferência de tipo
:=
fmt.Println
*/

func VariaveisETipos() {

	var name string
	var idade int
	var altura float64
	var estudando bool

	fmt.Println(name, idade, altura, estudando)

}

/*
Exercício 2 — Conversão de temperatura

Crie uma variável contendo uma temperatura em Celsius:

25

Converta para Fahrenheit usando:

F = C * 9/5 + 32

Imprima o resultado.

Desafio extra: faça o programa funcionar também com valores negativos.
*/

func ConversaoTemperatura() {
	var celsius float64 = 25
	var fahrenheit = celsius*9/5 + 32
	var negativeValue = celsius - fahrenheit

	fmt.Println(fahrenheit)
	fmt.Println(negativeValue)
}

/*
Exercício 3 — Troca de valores

Crie duas variáveis:

a := 10
b := 20

Troque os valores delas para que:

a = 20
b = 10

Regra: não crie uma terceira variável.

Esse exercício é simples, mas importante para entender uma característica interessante da sintaxe Go.
*/

func TrocaDeValores() {
	//a := 10
	//b := 20
	a, b := 10, 20
	fmt.Println(a, b, a)
}
