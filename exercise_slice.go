package main

import "fmt"

// Dado este slice de structs:

type Pedido struct {
	ID    int
	Valor float64
	Pago  bool
}

func exercise_slice() {
	pedidos := []Pedido{
		{ID: 1, Valor: 120.0, Pago: true},
		{ID: 2, Valor: 85.0, Pago: false},
		{ID: 3, Valor: 200.0, Pago: true},
		{ID: 4, Valor: 60.0, Pago: false},
	}

	total := 0.0
	for _, pedido := range pedidos {
		if !pedido.Pago {
			total += pedido.Valor
		}
	}
	fmt.Printf("Total não pago: R$ %.2f\n", total)
}

/*
Desafio: usando for range, calcule a soma total apenas dos pedidos não pagos e imprima o resultado.
*/

func exercise_slice2() {
	type Produto struct {
		Nome       string
		Quantidade int
		Preco      float64
	}

	produtos := []Produto{
		{Nome: "Teclado", Quantidade: 10, Preco: 150.0},
		{Nome: "Mouse", Quantidade: 3, Preco: 80.0},
		{Nome: "Monitor", Quantidade: 7, Preco: 900.0},
		{Nome: "Headset", Quantidade: 1, Preco: 200.0},
	}

	for _, produto := range produtos {
		if produto.Quantidade <= 3 {
			valorEmEstoque := float64(produto.Quantidade) * produto.Preco
			fmt.Printf("Estoque baixo: %s - quantidade: %d - valor em estoque: R$ %.2f\n",
				produto.Nome, produto.Quantidade, valorEmEstoque)
		}
	}
}

func exercice_slice_without_range() { // para apagar uma slice, cria-se outra.
	games := []string{"assasins creed", "ow", "fortnite"}
	wihout_For := games[:]
	fmt.Println("wihout_For: ", wihout_For)
}

// anexando uma slice
