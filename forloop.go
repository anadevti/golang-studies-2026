package main

import "fmt"

// 1. x := 10        → inicialização (1x)
// 2. x < 100?       → condição (verifica antes de cada iteração)
// 3. fmt.Println(x) → corpo do loop
// 4. x++            → pós (executa após cada iteração)
// 5. volta ao passo 2
func testLoop() {
	for x := 10; x < 100; x++ {
		fmt.Println(x)
	}
}

// Utilizando um while sem ter while em go:
func testLoopWhileCondition() {
	y := 1       // variável declarada FORA do for
	for y < 20 { // só a condição, igual ao while de outras linguagens
		fmt.Printf("%v ", y)
		y++ // incremento DENTRO do body da func
	}
}

// https://go.dev/doc/effective_go#for
func exampleFor() {

	// 1. For completo (como um for tradicional)
	for y := 1; y < 20; y++ {
	}

	// 2. Só a condição (simula o WHILE)
	y := 1
	for y < 20 {
		y++
	}

	// 3. Sem nada (loop infinito, simula o while true)
	for {
		break // usa break para sair
	}
}
