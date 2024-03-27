package main

import (
	"fmt"
	"reflect"
)

type Forma interface {
	Nombre() string
	calcularArea() float64
	calcularPerimetro() float64
}

type Circulo struct {
	radio float64
}

type Rectangulo struct {
	alto  float64
	ancho float64
}

type Triangulo struct {
	lado1 float64
	lado2 float64
	lado3 float64
}

func (c Circulo) Nombre() string {
	return reflect.TypeOf(c).Name()
}

func (c Circulo) calcularArea() float64 {
	return 3.14 * (c.radio * c.radio)
}

func (c Circulo) calcularPerimetro() float64 {
	return 3.14 * (c.radio * 2)
}

func (r Rectangulo) Nombre() string {
	return reflect.TypeOf(r).Name()
}

func (r Rectangulo) calcularArea() float64 {
	return r.alto * r.ancho
}

func (r Rectangulo) calcularPerimetro() float64 {
	return 2 * (r.alto + r.ancho)
}

func calcularAreaYPerimetro(f Forma) {
	fmt.Printf("Area del %s: %v ", f.Nombre(), f.calcularArea())
	fmt.Printf("Perimetro del %s: %v\n", f.Nombre(), f.calcularPerimetro())
}

func main() {
	c := Circulo{7}
	r := Rectangulo{5, 4}
	calcularAreaYPerimetro(c)
	calcularAreaYPerimetro(r)
}
