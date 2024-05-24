package main

import "fmt"

func main() {

	//ARRAYS
	/*
		Un array en Go es una estructura de datos que almacena una colección
		de elementos del mismo tipo.

		Los arrays en Go son estáticos, lo que significa que su tamaño no puede
		cambiarse una vez que se crean.

		Los Arrays en Go se declaran mediante la siguiente sintaxis:

		var nombreArray [longitud] tipoDato

		Donde:

			nombre_array es el nombre del array.
			longitud es el número de elementos que el array puede almacenar.
			tipo es el tipo de datos de los elementos del array.


		Por ejemplo, para declarar un array de 10 enteros, se utilizaría la siguiente sintaxis:

		var numeros [10]int

		Los elementos de un array se pueden acceder utilizando el operador de indexación []. Por ejemplo, para acceder al primer elemento del array numeros, se utilizaría la siguiente sintaxis:

		numeros[0]

		Los arrays en Go también se pueden inicializar al declararlos. Por ejemplo, para declarar un array de 10 enteros e inicializarlo con los valores 0, 1, 2, ..., 9, se utilizaría la siguiente sintaxis:

		var numeros [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}


		Los Arrays también se pueden declarar sin tamaño definido:

		var nombre_array = [...]tipo{elementos}

		También hay una manera más corta de declarar Arrays:

		nombre_array := [...]tipo{elementos}

		Para crear un array de forma corta vacío:

		nombre_array := make([]tipo, longitud)
	*/

	//EJEMPLO:

	// var n int
	// println("Ingrese el número de elementos a procesar: ")
	// fmt.Scanln(&n)

	// var numeros []int = make([]int, n)

	// for i := 0; i < n; i++ {
	// 	var input int
	// 	println("Ingrese un número para añadir a la lista: ")
	// 	fmt.Scanln(&input)
	// 	numeros[i] = input
	// }

	// fmt.Println(numeros)

	arrayStrings := []string{"Hola", "Mundo"}

	for _, elemento := range arrayStrings {
		fmt.Printf("Elemento: %s\n", elemento)
	}

	//SLICES
	/*
		Un slice en Go es una estructura de datos que representa una subsección de un array. Los slices son dinámicos, lo que significa que su tamaño puede cambiarse en tiempo de ejecución.

		Los slices se crean utilizando la siguiente sintaxis:

		nombre_slice := []tipo{valor1, valor2, ..., valorN}

		También se puede crear un slice sin especificar los valores iniciales. En este caso, el slice se inicializará con valores cero. Por ejemplo, para crear un slice de 10 enteros inicializados con valores cero, se utilizaría la siguiente sintaxis:

		numeros := make([]int, 10)

		Sin embargo, a diferencia de los arrays, los slices no tienen una longitud fija. Esto significa que se pueden añadir o eliminar elementos del slice en tiempo de ejecución.

		Por ejemplo, para añadir un elemento al slice numeros, se utilizaría la siguiente sintaxis:

		numeros = append(numeros, 10)

		La función append() añade un nuevo elemento al slice y devuelve el slice actualizado.

		Para eliminar un elemento del slice numeros, se utilizaría la siguiente sintaxis:

		numeros = append(numeros[:i], numeros[i+1:]...)

		Donde i es el índice del elemento que se desea eliminar.
	*/

	//EJEMPLO:

	calificaciones := make([]int, 4) //[0 0 0 0]
	var input int

	for i := 0; i < 4; i++ {
		fmt.Printf("Ingrese la calificación %d: ", i+1)
		fmt.Scanln(&input)
		calificaciones[i] = input
	}

	fmt.Println("Calificaciones: ", calificaciones)
	suma := suma_array(calificaciones)
	fmt.Println("Suma: ", suma)

}

//FUNCIONES

/*

	func nombre (parámetro tipo) tipo_de_retorno {
		//Cuerpo de la función
		return valor_a_retornar
	}
*/

func suma_array(numeros []int) int {
	var suma int = 0
	for _, elemento := range numeros {
		suma += elemento
	}
	return suma
}
