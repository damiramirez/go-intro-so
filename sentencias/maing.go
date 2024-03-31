package main

import "fmt"

func main() {
	var edad int = 76;
	if edad < 18 {
		fmt.Println("Menor de edad!")
	} else if(edad < 75) {
		fmt.Println("Mayor de edad!")
	} else {
		fmt.Println("Jubilado")
	}

	// Switch

	funcionCase("A")
	funcionCase("D")
	funcionCase("J")

	// For
	for i:=0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// While - No existe en go, es un for 
	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
}

func funcionCase(a string) {
	switch a {
	case "A":
		fmt.Println("Abeja")
	case "B":
		fmt.Println("Baskett")
	case "C":
		fmt.Println("Codigo")
	case "D":
		fmt.Println("Dado")
	default:
		fmt.Println("No se que es!")
	}
}