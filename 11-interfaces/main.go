package main

import "fmt"

func main() {
	params := Params{name: "lucas", age: 28}
	c := creadoraDeMember{}
	member := creadoraCaller(c, params)
	imprimirCaller(member)
	gritarCaller(member)
	paramsParent := Params{name: "pariente", dni: "38749898"}
	cp := creadoraDeParent{}
	parent := creadoraCaller(cp, paramsParent)
	imprimirCaller(parent)
	gritarCaller(parent)
}

type IModel interface {
	imprimir()
	gritar()
}

type Params struct {
	name string
	age  int
	dni  string
}

type Member struct {
	name string
	age  int
}

type Parent struct {
	name string
	dni  string
}

func (m Member) imprimir() {
	fmt.Println(m)
}

func (m Member) gritar() {
	fmt.Println("AAAAAAAAAAA")
}

func (m Parent) imprimir() {
	fmt.Println(m, "PARIENTE")
}

func (m Parent) gritar() {
	fmt.Println("EEEEEEEEEEEEEEEEEEEE")
}

type creadoraDeModelos interface {
	crearModelo(Params) IModel
}

type creadoraDeMember struct{}

func (c creadoraDeMember) crearModelo(params Params) IModel {
	member := Member{}
	member.name = params.name
	member.age = params.age
	return member
}

type creadoraDeParent struct{}

func (c creadoraDeParent) crearModelo(params Params) IModel {
	parent := Parent{}
	parent.name = params.name
	parent.dni = params.dni
	return parent
}

func creadoraCaller(c creadoraDeModelos, p Params) IModel {
	return c.crearModelo(p)
}

func gritarCaller(m IModel) {
	m.gritar()
}

func imprimirCaller(m IModel) {
	m.imprimir()
}
