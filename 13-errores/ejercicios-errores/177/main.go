package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Fatalln("ha ocurrido un error:", err)
	}
	fmt.Println(string(bs))

}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error en aJSON: %v", err)
		// return []byte{}, errors.New(fmt.Sprintf("ha ocurrido un error: %v", err))
	}
	fmt.Println(bs)
	return bs, nil
}

// fmt.Sprintf() devuelve un string
// fmt.Errorf() devuelve un error
