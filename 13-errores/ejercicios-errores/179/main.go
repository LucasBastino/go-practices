package main

import (
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("%v error matemático: %v %v", re.err, re.lat, re.long)
	// devuelve el string
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err.Error())
		// log.Println(err)
		// esto es lo mismo pero a mi no me parece tan claro
		// especificando err.Error() me parece mas entendible
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		error1 := fmt.Errorf("ha ocurrido un error,")
		return 0, raizError{"50.2289 N", "99.4656 W", error1}
	}
	return 42, nil
}
