package main

import (
	"fmt"
	"sync"
)

var valor int = 0
var wg sync.WaitGroup
// var mutex sync.Mutex
var sem = make(chan int, 1)

func main() {

	for i := 0; i < 20; i++ {
		// Agrego uno al WaitGroup antes de comenzar cada rutina
		wg.Add(1)
		go thread()
	}

	// Esperar a que todas las rutinas finalicen
	wg.Wait()
}

func thread() {
	// Indico que al finalizar la función thread tiene que liberar un valor de wg
	defer wg.Done()

	//Obtenemos el candado de mutex
	// mutex.Lock()

	// representar un mutex con un canal
	sem <- 0
	valor++
	fmt.Println("Este es el hilo número", valor)

	//Soltamos el candado de mutex
	// mutex.Unlock()

	<-sem
}