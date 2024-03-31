package main

import "fmt"

var valor1 int = 3

func main() {
	var valor1 int = 1
	var valor2 int = 2

	// Scopes
	// Golang primero va a buscar la variable desde el scope actual hasta el mas externo
	
	// Usa scope de la funcion de valor1
	fmt.Println("Resultado sumar:", sumar(valor1, valor2))
	// Usa scope global de valor1
	fmt.Println("Resultado sumar2:", sumar2(valor2))
}

// func nombreFuncion(valor tipoDato) tipoDatoReturn

func sumar(a int, b int) int {
	return a + b
}

func sumar2(b int) int {
	return valor1 + b
}