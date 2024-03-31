package main

import (
	"fmt"
	"time"
)

var valor int = 0

func main() {

	for i := 0; i < 20; i++ {
		go thread()
	}

	//Este tiempo es importante debido a que si el hilo
	//principal termina antes, los hilos no se
	//ejecutarán
	time.Sleep(101 * time.Second)
}

func thread() {
	valor++
	fmt.Println("Este es el hilo número", valor)
}