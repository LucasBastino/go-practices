package main

import "fmt"

type Parcial struct {
	materia string
	nota    int
}

type Materia struct {
	nombre    string
	parciales []Parcial
	// final     int
}

type Alumno struct {
	nombre   string
	edad     int
	materias []Materia
}

func main() {

	crearDatos()
}

func crearDatos() {
	parcial1 := Parcial{"Matematica", 2}
	fmt.Println(parcial1)
	materia1 := Materia{"Matematica", []Parcial{parcial1}}
	fmt.Println(materia1)
	// alumno1 := Alumno("Lucas", 27, [materia1, materia2])
}
