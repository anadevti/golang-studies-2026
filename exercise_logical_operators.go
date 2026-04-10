package main

import (
	"fmt"
	"time"
)

// Operadores lógicos: && (AND), || (OR), ! (NOT)
func exerciseLogicalOperators() {
	//true && true -> true
	//true && false -> true
	//true || true -> true
	//true || false -> false
	//!true -> false
}

func exerciseLogicalOperators2() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exerciseLogicalOperators3() {
	// 65 a 90
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("%\t%#U\n ", i)
		}
	}
}

func exerciseLogicalOperators4() {
	// O for condition não tem init nem post

	year := time.Date(2001, time.May, 24, 0, 0, 0, 0, time.UTC)
	for year.Before(time.Now()) {
		fmt.Println("Os Anos desde que eu nasci sao:", year.Year())
		year = year.AddDate(1, 0, 0)
	}
}

func exerciseLogicalOperators5() {
	// aqui temos inicialização, condição e pós

	for i := 2001; i <= 2026; i++ {
		fmt.Println(i)
	}
}
