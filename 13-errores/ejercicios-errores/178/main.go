package main

import "fmt"

func main() {
	err := errorPer{"no se puede dividir por cero", 2}
	foo(err)
	foo2(err)
}

type errorPer struct {
	info   string
	status int
}

func (e errorPer) Error() string {
	return fmt.Sprintf("Ha ocurrido un error: %s\nStatus: %v", e.info, e.status)
}

// implementar la func Error hace que sea de tipo error y ademas que devuelva un string
// cuando se imprime la variable, no imprime los campos

func foo(err error) {
	fmt.Println(err)

}

func foo2(err error) {
	fmt.Println(err.(errorPer).info)
	// este assertion se hace cuando estas 100% seguro que la variable es de ese tipo
}
