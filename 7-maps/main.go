package main

import "fmt"

func main() {
	fmt.Println("maps exercises:")
	// crearMapas()
	// cambiarElementosEnMapas()
	leerElementosEnMapas()
}

// los mapas son como mini bases de datos

// func crearMapas() {
// 	// se puede inicializar vacio
// 	map1 := map[int]string{}

// 	// y tambien con algunos elementos
// 	map2 := map[int]string{
// 		32788265: "Maxi",
// 		41259869: "Camila",
// 		35658478: "Ana",
// 	}

// 	fmt.Println("map1:", map1)
// 	fmt.Println("map2:", map2)
// }

// func cambiarElementosEnMapas() {
// 	map1 := map[int]string{
// 		32788265: "Maxi",
// 		41259869: "Camila",
// 		35658478: "Ana",
// 	}

// 	map1[45582316] = "Lisandro" // agrego un elemento
// 	map1[41259869] = "Camila2"  // reemplazo un elemento

// 	fmt.Println(map1)

// }

func leerElementosEnMapas() {
	map1 := map[int]string{
		32788265: "Maxi",
		41259869: "Camila",
		35658478: "Ana",
		44886396: "",
	}
	dato1 := map1[32788265]

	dato2, existe := map1[44886396]
	// dato2, existe := map1[11111111]
	// si el dato esta vacio o no existe devolvera el valor cero
	// por lo que para saber si esta vacio requiere de un segundo argumento

	if existe {
		fmt.Println("el elemento existe y esta vacio")
	} else {
		fmt.Println("el elemento no existe")
	}
	fmt.Println("dato1:", dato1)
	fmt.Println("dato2:", dato2)

}
