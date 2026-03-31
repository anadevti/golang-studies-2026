package main

import "fmt"

func calcularFrete(pesoKg float64, distanciaKm float64) (float64, error) {
	// 1. validar entradas
	if pesoKg <= 0 {
		return 0, fmt.Errorf("peso inválido: %.2f", pesoKg)
	}
	if distanciaKm <= 0 {
		return 0, fmt.Errorf("distancia inválida: %.2f", distanciaKm)
	}

	// 2. determinar preço por kg
	var precoPorKg float64

	if pesoKg <= 5 {
		precoPorKg = 2.00
	} else if pesoKg <= 20 {
		precoPorKg = 1.50
	} else {
		precoPorKg = 1.00
	}

	// 3. determinar multiplicador de distância
	var multiplicador float64

	if distanciaKm <= 100 {
		multiplicador = 1.0
	} else if distanciaKm <= 500 {
		multiplicador = 1.5
	} else {
		multiplicador = 2.0
	}

	// 4. calcular e retornar
	frete := pesoKg * precoPorKg * multiplicador
	return frete, nil
}
