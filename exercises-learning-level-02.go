package main

import (
	"fmt"
)

// %d - decimal
// %b - binario
// %#x - hexadecimal
func Exercise2() {
	number := 20

	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binario: %b\n", number)
	fmt.Printf("Hexadecimal: %#x\n", number)
}

func Exercise3() {
	a := 10
	b := 20
	c := 30
	d := 40

	if a > 18 {
		fmt.Println("maior de idade")
	} else {
		fmt.Println("menor de idade")
	}

	if b == 20 {
		fmt.Println("b é igual a 20")
	}

	if c >= 20 {
		fmt.Println("c é maior ou igual a 20")
	}

	if c != d {
		fmt.Println("c é diferente de d")
	}

	if d <= 50 {
		fmt.Println("d é menor ou igual que 40")
	}
	fmt.Println("*****************************")
	fmt.Printf("a é do tipo %T\n b é do tipo: %T\n c é do tipo: %T\n d é do tipo: %T\n", a, b, c, d)
}

func Exercise4() {
	const name int = 10 // tipado
	const test = 20     // nao tipado

	fmt.Printf("name é %v Test é %v", name, test)
}

func Exercise5() {
	var test = 20
	fmt.Printf("Decimal: %d Binario: %b, Hexadecimal: %#x\n", test, test, test)
	x := test << 1 // deslocando bits :0
	fmt.Printf("Decimal: %d Binario: %b, Hexadecimal: %#x", x, x, x)
}

func Exercise6() {
	var rawString = `aninha do javinha.`
	fmt.Println(rawString)
}

func Exercise7() {
	const (
		_ = 2026 + iota
		b
		c
		d
		e
	)
	fmt.Printf("B possui o valor: %v C possui o valor: %v, D possui o valor: %v, E possui o valor: %v", b, c, d, e)
}
