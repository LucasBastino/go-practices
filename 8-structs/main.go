package main

import (
	f "fmt"

	"github.com/LucasBastino/practicas-go/sliceMethods"
)

func main() {
	f.Println("structs exercises:")
	// crearStruct()
	// borrarDuplicados()
	sliceMethods.Mostrar2()
}

// func crearStruct() {
// 	type jugador struct {
// 		nombre   string
// 		edad     int
// 		posicion string
// 		salario  int
// 	}

// 	jugador1 := jugador{"alustiza", 33, "mediocampista", 5000}
// 	f.Println(jugador1.nombre)
// 	jugador1.posicion = "delantero"
// 	f.Println(jugador1.posicion)
// 	f.Println("salario en pesos:", jugador1.salario*1100)

// 	f.Printf("%v\n", jugador1)
// 	f.Printf("%+v\n", jugador1)
// 	f.Printf("%#v\n", jugador1)
// }

// func borrarDuplicados() {
// 	jugadores := []string{
// 		"correa", "alustiza", "perdomo", "godoy", "russo", "casa", "russo", "godoy",
// 		"gimenez", "alustiza", "achucarro", "perdomo", "gimenez", "zarif", "vismara",
// 		"correa", "godoy",
// 	}

// 	// creo un mapa vacio con clave string y valor struct vacio
// 	unicos := map[string]struct{}{}

// 	// recorro el slice de duplicados y agrego cada jugador como clave de un mapa
// 	// si uno ya se encuentra, lo sobreescribe, por lo tanto
// 	// quedan los jugadores como clave, y los duplicados ya fueron sobreescritos
// 	for _, jugador := range jugadores {
// 		unicos[jugador] = struct{}{}
// 	}

// 	// el for con maps es "for key, value := range... "
// 	for jugador := range unicos {
// 		f.Println(jugador)
// 	}

// }
