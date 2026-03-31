package main

//Escreva uma função chamada classificarTemperatura que recebe um float64 e retorna uma string com a classificação:
//
//abaixo de 0: "congelante"
//entre 0 e 15: "frio"
//entre 15 e 25: "agradável"
//acima de 25: "quente"

func temperatureClass(temp float64) string { // utilizando early return para evitar aninhamento
	if temp <= 0 {
		return "congelante"
	}
	if temp <= 15 {
		return "frio"
	}
	if temp <= 25 {
		return "agradável"
	}
	return "quente"
}
