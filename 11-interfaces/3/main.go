package main

import "fmt"

func main() {
	// params := Params{model: "member", name: "lucas", age: 28}
	params := Params{model: "parent", name: "roman", dni: 22552255}
	model := crearModel(params)
	imprimirModel(model)

}

type IModel interface {
	imprimirse()
}

type Params struct {
	model string
	name  string
	age   int
	dni   int
}

type Member struct {
	name string
	age  int
}

type Parent struct {
	name string
	dni  int
}

func crearModel(params Params) IModel {
	switch params.model {
	case "member":
		return crearMember(params)
	case "parent":
		return crearParent(params)
	default:
		return nil
	}
}

func crearMember(params Params) IModel {
	return Member{params.name, params.age}

}

func crearParent(params Params) IModel {
	return Parent{params.name, params.dni}

}

func (m Member) imprimirse() {
	fmt.Printf("este es un member y su nombre es %s, su edad %v\n", m.name, m.age)
}

func (p Parent) imprimirse() {
	fmt.Printf("este es un parent y su nombre es %s, su dni %v\n", p.name, p.dni)
}

func imprimirModel(model IModel) {
	model.imprimirse()
}
