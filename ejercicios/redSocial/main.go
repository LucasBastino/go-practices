package main

import "fmt"

type Usuario struct {
	nombre  string
	edad    int
	amigos  []Usuario
	posteos []string
	activo  bool
}

func main() {
	lucas := Usuario{"Lucas", 22, []Usuario{}, []string{}, true}
	lucas.imprimirInfo()
}

func (u Usuario) imprimirInfo() {
	fmt.Println("Nombre de usuario:", u.nombre)
	fmt.Println("Edad:", u.edad)
	fmt.Println("Lista de amigos:", u.amigos)
	if u.activo {
		fmt.Println("Esta activo? Si")
	} else {
		fmt.Println("Esta activo? No")
	}
}

func (u *Usuario) modificarAtributo(atributo string, valor any) {
	fmt.Println(u.atributo)
}
