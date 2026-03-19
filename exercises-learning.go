package main

import (
	"fmt"
)

func exerciseLearning() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exerciseExerciseTwo() {

	var x = 42
	var y = "James Bond"
	var z = true

	s := fmt.Sprintf("%v,%v,%v", x, y, z)
	fmt.Println(s)
}

func exerciseType() {
	type test int
	var x test
	fmt.Printf("O tipo eh %T, e o valor '%v'", x, x)
	x = 42
	fmt.Println(x)
}
