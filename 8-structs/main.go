package main

import "fmt"

func main() {
	fmt.Println("structs exercises:")
	// crearStruct()
	borrarDuplicados()
}

// func crearStruct() {
// 	type jugador struct {
// 		nombre   string
// 		edad     int
// 		posicion string
// 		salario  int
// 	}

// 	jugador1 := jugador{"alustiza", 33, "mediocampista", 5000}
// 	fmt.Println(jugador1.nombre)
// 	jugador1.posicion = "delantero"
// 	fmt.Println(jugador1.posicion)
// 	fmt.Println("salario en pesos:", jugador1.salario*1100)

// 	fmt.Printf("%v\n", jugador1)
// 	fmt.Printf("%+v\n", jugador1)
// 	fmt.Printf("%#v\n", jugador1)
// }

func borrarDuplicados() {
	jugadores := []string{
		"correa", "alustiza", "perdomo", "godoy", "russo", "casa", "russo", "godoy",
		"gimenez", "alustiza", "achucarro", "perdomo", "gimenez", "zarif", "vismara",
		"correa", "godoy",
	}

	// creo un mapa vacio con clave string y valor struct vacio
	unicos := map[string]struct{}{}

	// recorro el slice de duplicados y agrego cada jugador como clave de un mapa
	// si uno ya se encuentra, lo sobreescribe, por lo tanto
	// quedan los jugadores como clave, y los duplicados ya fueron sobreescritos
	for _, jugador := range jugadores {
		unicos[jugador] = struct{}{}
	}

	// el for con maps es "for key, value := range... "
	for jugador := range unicos {
		fmt.Println(jugador)
	}

}
