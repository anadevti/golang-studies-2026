package praticas

import (
	"fmt"
	"strings"
)

func testMap() {
	friends := map[string]int{
		"John": 40404040,
		"Jane": 41414141,
	}

	friends["Gio"] = 98989898 // adicionando valor a um map

	fmt.Println(friends)
}

// Escreva uma função que recebe uma string e retorna um
// map[string]int com a contagem de quantas vezes cada palavra aparece.
func countWords(s string) map[string]int {
	words := strings.Fields(s)      // Divide a string por espaços em branco
	mapKeys := make(map[string]int) // inicializando o map

	for _, word := range words {
		mapKeys[word] = mapKeys[word] + 1
	}

	fmt.Println(mapKeys)
	return mapKeys
}

// Você tem uma slice de transações. Agrupe-as num map[string][]Transaction,
// onde a chave é o tipo da transação ("credito", "debito", "pix").
type Transaction struct {
	ID    string
	Type  string
	Value float64
}

func GroupByType(t []Transaction) map[string][]Transaction {
	grouped := make(map[string][]Transaction)

	for _, transaction := range t {
		grouped[transaction.Type] = append(grouped[transaction.Type], transaction)
	}

	return grouped
}

func SumTotalByType(grouped map[string][]Transaction) map[string]float64 {
	totals := make(map[string]float64) // map de saída: tipo -> soma

	for tipo, transacoes := range grouped {
		total := 0.0
		for _, transaction := range transacoes {
			total += transaction.Value
		}
		totals[tipo] = total
	}

	return totals
}
