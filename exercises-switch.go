package main

import (
	"fmt"
)

// Desafio 1 — Básico: Classificador de notas
//
// Escreva um programa que receba uma nota (inteiro de 0 a 100)
// e imprima a classificação correspondente:
//
// Nota         Classificação
// 90–100       Excelente
// 70–89        Aprovado
// 50–69        Recuperação
// 0–49         Reprovado
// Fora do intervalo  Nota inválida
//
// Requisitos:
// - Use switch (não if/else)
// - A lógica de faixa deve estar dentro do switch
// - O programa deve compilar e rodar com go run

// primeira resolução:
func classificationNotes() {
	note := 100

	switch note {
	case 90, 100:
		fmt.Println("Nota Excelente!")
	case 70, 89:
		fmt.Println("Aprovado!")
	case 50, 69:
		fmt.Println("ixi, voce esta de recuperação!")
	case 0, 49:
		fmt.Println("F, reprovado.")
	default:
		fmt.Println("Nota invalida.")
	}
}

func classificationNotesV2() {
	note := 100

	switch {
	case note >= 90 && note <= 100:
		fmt.Println("Nota Excelente!")
	case note >= 70 && note <= 89:
		fmt.Println("Aprovado!")
	case note >= 50 && note <= 69:
		fmt.Println("ixi, voce esta de recuperação!")
	case note >= 0 && note <= 49:
		fmt.Println("F, reprovado.")
	default:
		fmt.Println("Nota invalida.")
	}
}
