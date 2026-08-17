package praticas

import "fmt"

/*
Nível 2 — Condicionais
Exercício 4 — Par ou ímpar

Receba um número inteiro e informe:

10 -> par
7 -> ímpar

Use if.

Desafio extra: trate também o número 0.
*/

func Numbers(number int) {
	if number%2 == 0 {
		fmt.Println("par")
	} else {
		fmt.Println("ímpar")
	}

	Numbers(10)
	Numbers(7)
	Numbers(42)
	Numbers(15)
}

/*
Exercício 5 — Maior de três números

Dado:

a := 10
b := 35
c := 20

Descubra qual é o maior.

Não use funções prontas como max.
*/

func NumberMaior() {
	a, b, c := 10, 35, 20
	var maior int
	if a >= b && b >= c {
		maior = a
	} else if b >= a && b >= c {
		maior = b
	} else {
		maior = c
	}
	fmt.Printf("O maior número é: %d\n", maior)
}

/*
Exercício 6 — Classificação de idade

Crie uma variável idade e classifique:

0 - 12    -> criança
13 - 17   -> adolescente
18 - 59   -> adulto
60+       -> idoso

Tente escrever de uma maneira que seja fácil de ler.
*/
func Ages(idade int) {
	if idade <= 12 {
		fmt.Println("voce ainda é uma crianca")
	} else if idade <= 17 {
		fmt.Println("voce ainda é um adolescente")
	} else if idade <= 59 {
		fmt.Println("voce ainda é um adulto")
	} else if idade >= 60 {
		fmt.Println("voce ainda é um idoso")
	}
}

/*
Exercício 7 — switch

Crie uma variável:

dia := 3

Use switch para imprimir:

1 -> Domingo
2 -> Segunda
3 -> Terça
...
7 -> Sábado

Caso o número seja inválido, imprima:

Dia inválido
*/
func dias() {
	dia := 3
	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terca")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("Dia inválido")
	}
}

/*
Use for para imprimir de 1 até 10.
Depois, modifique para imprimir apenas os números pares.
*/

func TestFor() {
	// inicializacao // condicao // pos
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}
}

/*
Calcule a soma dos números de 1 até 100 sem colocar o resultado
diretamente no código.
Resultado esperado: 5050.
*/
func testFor2() {
	sum := 0
	for j := 1; j <= 100; j++ {
		sum += j
	}
	fmt.Println("A soma de 1 até 100 é:", sum)
}

/*
### Exercício 10 — Tabuada

Dado `numero := 7`, imprima a tabuada de `1` a `10`.
*/

func testFor3() {
	num := 7
	for j := 1; j <= 10; j++ {
		i := num * j
		fmt.Println("A tabuada de 1 até 10 é:", i)
	}

}
