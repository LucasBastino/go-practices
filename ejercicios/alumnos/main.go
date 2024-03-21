package main

import "fmt"

type Alumno struct {
	nombre string
	edad   int
	notas  map[string][]int
}

func main() {

	alumno1 := Alumno{"Lucas", 27, map[string][]int{}}
	alumno1.notas["Matematica"] = []int{5, 8, 7, 9}
	alumno1.notas["Geografia"] = []int{4, 10, 5, 8}
	fmt.Println(alumno1)
	imprimirParciales(alumno1)
	imprimirTodosLosParciales(alumno1)
	sumarleNota(alumno1)
}

func imprimirParciales(alumno Alumno) {
	fmt.Println(alumno.nombre, alumno.edad)
	for key, value := range alumno.notas {
		fmt.Println(key, value)
	}
	fmt.Println(alumno.notas)
}

func imprimirTodosLosParciales(alumno Alumno) {
	fmt.Println(alumno.nombre, alumno.edad)
	for key, value := range alumno.notas {
		fmt.Println(key)
		for i := 0; i < len(value); i++ {
			fmt.Println(value[i])
		}
	}
}

func sumarleNota(alumno Alumno) {
	fmt.Println("sumarle nota")
	for _, value := range alumno.notas {
		for i, nota := range value {
			nota++
			fmt.Println(nota)
			fmt.Println(value[i])
		}
	}
	fmt.Println(alumno.notas)
}
