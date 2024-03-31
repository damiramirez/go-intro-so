package main

import "fmt"

func main() {

	// TIPOS DE DATOS BASICOS
	// Variables sin valor -> defualt
	// - int, float32, byte-> 0
	// - bool -> false
	// - string -> ""

	// Integer
	noTipado := 123456
	var tipadoSinValor int
	var tipadoConValor int = 456789
	fmt.Println("Valor no tipado: ", noTipado)
	fmt.Println("Valor tipado sin valor: ", tipadoSinValor)
	fmt.Println("Valor tipado con valor: ", tipadoConValor)

	// Bool
	var booleanSinValor bool
	var booleanConValor bool = true
	fmt.Println("Boolean sin valor: ", booleanSinValor)
	fmt.Println("Boolean con valor: ", booleanConValor)

	// String
	var stringSinValor string
	var stringConValor string = "Soy una cadena"
	fmt.Println("String sin valor: ", stringSinValor)
	fmt.Println("String con valor: ", stringConValor)

	// Float32 - Float64
	var floatSinValor float32
	var floatConValor float32 = 17.23
	fmt.Println("Flotante sin valor: ", floatSinValor)
	fmt.Println("Flotante con valor: ", floatConValor)

	// Byte
	var byteSinValor byte
	var byteConValor byte = 1
	fmt.Println("Byte sin valor: ", byteSinValor)
	fmt.Println("Byte con valor: ", byteConValor)

	// TIPOS DE DATOS COMPUESTOS

	// Array - Longitud fija
	var arraySinValor [5]int
	var arrayConValor [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array sin valor: ", arraySinValor)
	fmt.Println("Array con valor: ", arrayConValor)
	fmt.Println("Valor 3 del array: ", arrayConValor[2])

	// Slice - Longitud dinamica
	// - append(slice, elem...) slice
	// - Los slice no tienen para remover x indice- debo crear subslices

	var sliceSinValor []int
	var sliceConValor []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice sin valor: ", sliceSinValor)
	fmt.Println("Slice con valor: ", sliceConValor)
	var tamañoSliceConValor = len(sliceConValor)
	fmt.Println("Tamaño del slice con valor: ", tamañoSliceConValor)
	sliceConValor = append(sliceConValor, 6)
	tamañoSliceConValor = len(sliceConValor)
	fmt.Println("Nuevo tamaño del slice con valor: ", tamañoSliceConValor)
	fmt.Println("Nuevo slice con valor: ", sliceConValor)
	fmt.Println("Valor 3 del slice: ", sliceConValor[2])

	// Map - [key]value

	var mapSinValor map[string]int
	var mapConValor map[string]int = map[string]int{"Juan": 28, "Ana": 30}
	fmt.Println("Map sin valor: ", mapSinValor)
	fmt.Println("Map con valor: ", mapConValor)
	fmt.Println("Valor para Juan: ", mapConValor["Juan"])
	fmt.Println("Valor para NoExiste: ", mapConValor["NoExiste"])
	mapConValor["Nahuel"] = 32
	fmt.Println("Valor para Nahuel: ", mapConValor["Nahuel"])
	delete(mapConValor, "Ana")
	fmt.Println("Map con valor: ", mapConValor)

	// Struct - Conjunto de campos con difrentes tipos de datos

	type Persona struct {
		nombre string
		edad   int
		altura float32
	}

	var structSinValor Persona
	var structConValor Persona = Persona{"Nahuel", 32, 1.77}
	fmt.Println("Struct sin valor:", structSinValor)
	fmt.Println("Struct con valor:", structConValor)
	fmt.Println("Nombre del struct:", structConValor.nombre)
	structConValor.edad = 33
	fmt.Println("Edad del struct:", structConValor.edad)

	// String a byte - Util para el TP
	// - No siempre una secuencia de N caracteres va a ocupar N bytes

	var cadena string = "Hola"
	var arrayBytes []byte
	arrayBytes = []byte(cadena)
	fmt.Println("Array de bytes:", arrayBytes)
	var nuevaCadena string = string(arrayBytes)
	fmt.Println("Nueva cadena:", nuevaCadena)
}
