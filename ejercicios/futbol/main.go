package main

import (
	"fmt"

	"github.com/LucasBastino/practicas-go/sliceMethods"
)

type Equipo struct {
	nombre         string
	fundacion      int
	presupuestoEnM float64
	jugadores      []Jugador
}

type Jugador struct {
	nombre     string
	edad       int
	posicion   string
	valoracion int
	salario    float64
	valorEnM   float64
}

func main() {
	fmt.Println("Futbol")
	oroz := Jugador{"Oroz", 25, "MCO", 85, 5000, 2}
	vivaldo := Jugador{"Vivaldo", 39, "ARQ", 82, 3000, 1.5}
	// fmt.Println(oroz)
	chacarita := Equipo{"Chacarita Jrs", 1906, 20, []Jugador{oroz, vivaldo}}
	// fmt.Println(chacarita)
	sliceMethods.ShowField(chacarita.jugadores, "nombre")
	chacarita.eliminarJugador("Vivaldo")
	// fmt.Println(chacarita)

}

func (e *Equipo) eliminarJugador(value any) {
	e.jugadores = sliceMethods.DeleteOneByField(e.jugadores, "nombre", value)
	// fmt.Println(e.jugadores)
}
