package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Alumno struct {
	nombre         string
	edad           float64
	calificaciones map[string][]float64
}

func main() {

	alumno1 := Alumno{"Lucas", 27, map[string][]float64{}}
	alumno1.calificaciones["Matematica"] = []float64{}
	alumno1.calificaciones["Geografia"] = []float64{}
	alumno1.calificaciones["Lengua"] = []float64{}
	// imprimirParciales(alumno1)
	// imprimirTodosLosParciales(alumno1)
	// sumarleNota(alumno1)
	rendirParciales(alumno1)
	// fmt.Println(alumno1.calificaciones)
	pasoDeAño(alumno1)

}

func imprimirParciales(alumno Alumno) {
	fmt.Println(alumno.nombre, alumno.edad)
	for key, value := range alumno.calificaciones {
		fmt.Println(key, value)
	}
	fmt.Println(alumno.calificaciones)
}

func imprimirTodosLosParciales(alumno Alumno) {
	fmt.Println(alumno.nombre, alumno.edad)
	for key, value := range alumno.calificaciones {
		fmt.Println(key)
		for i := 0; i < len(value); i++ {
			fmt.Println(value[i])
		}
	}
}

func sumarleNota(alumno Alumno) {
	fmt.Println("sumarle nota")
	for _, value := range alumno.calificaciones {
		for i := range value {
			value[i]++
		}
	}
	fmt.Println(alumno.calificaciones)
}

// este no funciona, modifica una copia pero no el original, estan guardados en diferentes direcciones de memoria
// func sumarleNota(alumno Alumno) {
// 	fmt.Println("sumarle nota")
// 	for _, value := range alumno.calificaciones {
// 		for _, v := range value {
// 			v++
// 		}
// 	}
// 	fmt.Println(alumno.calificaciones)
// }

func rendirParcial() float64 {
	// mathRound(x*100)/100 -> hace que tu numero tenga 2 decimales
	nota := math.Round((rand.Float64()*10+3)*100) / 100
	if nota > 10 {
		nota = 10
	}
	return nota
}

func rendirParciales(alumno Alumno) {
	for materia := range alumno.calificaciones {
		for i := 0; i < 3; i++ {
			alumno.calificaciones[materia] = append(alumno.calificaciones[materia], rendirParcial())
			// fmt.Println(alumno.calificaciones[materia])
		}
	}
}

func pasoDeAño(alumno Alumno) {

	// promedio := 0
	for materia, parciales := range alumno.calificaciones {
		var suma float64 = 0
		for i := range parciales {
			suma += parciales[i]
			suma = math.Round(suma*100) / 100
		}
		fmt.Printf("%v: %v %v\n", materia, suma, parciales)
		promedio := math.Round(suma/float64(len(parciales))*100) / 100
		fmt.Println(promedio)
		if promedio >= 7 {
			fmt.Print("APROBO\n\n")
		} else {
			fmt.Print("DESAPROBO\n\n")
		}
	}
}
