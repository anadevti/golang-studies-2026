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
