package main

import "fmt"

func main() {
	c := creadoraDeMember{}
	params := Params{"lucas", 28}
	member := creadorCaller(c, params)
	callerImprimir(member)
	callerGritar(member)

}

type IModel interface {
	imprimirse()
	gritar()
}

type Member struct {
	name string
	age  int
}

func (m Member) imprimirse() {
	fmt.Println(m)
}

func (m Member) gritar() {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
}

type Params struct {
	name string
	age  int
}

type creadoraDeModelos interface {
	crearModelo(Params) IModel
}

type creadoraDeMember struct{}

func (c creadoraDeMember) crearModelo(params Params) IModel {
	var member Member
	member.name = params.name
	member.age = params.age

	return member
}

func creadorCaller(c creadoraDeModelos, params Params) IModel {
	return c.crearModelo(params)
}

func callerImprimir(m IModel) {
	m.imprimirse()
}

func callerGritar(m IModel) {
	m.gritar()
}
