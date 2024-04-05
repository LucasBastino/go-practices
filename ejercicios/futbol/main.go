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
	lesionado  bool
}

func main() {
	fmt.Println("Futbol")
	oroz := Jugador{"Oroz", 25, "MED", 85, 5000, 2, false}
	vivaldo := Jugador{"Vivaldo", 39, "ARQ", 82, 3000, 1.5, true}
	malcorra := Jugador{"Malcorra", 32, "MED", 86, 6000, 3, false}
	campaz := Jugador{"Campaz", 24, "DEL", 85, 5500, 2.5, false}
	chacarita := Equipo{"Chacarita Jrs", 1906, 20, []Jugador{oroz, vivaldo}}
	central := Equipo{"Rosario Central", 1889, 30, []Jugador{campaz, malcorra}}
	// fmt.Println(chacarita)
	// fmt.Println(central)
	chacarita.comprarJugador("Campaz", &central, 2)
	chacarita.imprimirInfo()
	central.imprimirInfo()
}

// func (e *Equipo) eliminarJugador(field string, value any) {
// 	e.jugadores = sliceMethods.DeleteOneByField(e.jugadores, field, value)
// 	// fmt.Println(e.jugadores)
// }

func (e *Equipo) comprarJugador(jugador string, equipoVendedor *Equipo, valor float64) {
	i := sliceMethods.FindIndexByField(equipoVendedor.jugadores, "nombre", jugador)
	e.jugadores = append(e.jugadores, equipoVendedor.jugadores[i])
	e.presupuestoEnM -= valor
	equipoVendedor.jugadores = sliceMethods.DeleteOneByIndex(equipoVendedor.jugadores, i)
	equipoVendedor.presupuestoEnM += valor

}

func (e Equipo) imprimirInfo() {
	fmt.Println(e.nombre)
	fmt.Println("Presupuesto:", e.presupuestoEnM)
	for _, jugador := range e.jugadores {
		fmt.Println(jugador.nombre)
	}
}
