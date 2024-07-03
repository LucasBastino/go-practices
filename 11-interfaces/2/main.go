package main

import "fmt"

func main() {
	a := Afiliado{"lucas", 3232323}
	f := Familiar{"prima", 25}
	imprimirModelo(a)
	imprimirModelo(f)
}

type InterfazModelo interface {
	imprimir()
}

type Afiliado struct {
	nombre string
	dni    int
}

type Familiar struct {
	nombre string
	edad   int
}

func imprimirModelo(modelo InterfazModelo) {
	modelo.imprimir()
}

func (afiliado Afiliado) imprimir() {
	fmt.Println(afiliado.nombre, afiliado.dni)
}

func (familiar Familiar) imprimir() {
	fmt.Println("este es un familiar y se dobla la edad", familiar.nombre, familiar.edad*2)
}
