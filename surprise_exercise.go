package main

import "fmt"

func surprise() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d, %#x \t \n  ", x, string(x)) // verbs
	}
}
