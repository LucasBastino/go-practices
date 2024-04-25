// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Name     string `json:"nombre"`
// 	Age      int
// 	Pass     bool
// 	Nickname string `json:"nick"`
// }

// func (s Student) String() string { // student.String() es lo mismo que fmt.Println(student1)
// 	return fmt.Sprintf("nam1111e: %s, age: %d, pass: %v, nickname: %s", s.Name, s.Age, s.Pass, s.Nickname)
// }

// func main() {
// 	studentInJson := `{"nombre":"Lucas","Age":20,"Pass":true,"nick":"fune"}`
// 	var studentInString Student
// 	err := json.Unmarshal([]byte(studentInJson), &studentInString)
// 	if err != nil {
// 		fmt.Println("error unmarshaling student")
// 		panic(err.Error())
// 	}
// 	fmt.Println(studentInString)
// 	fmt.Println(studentInString.String())
// 	// es lo mismo. Yo igualmente prefiero ser claro y ponerle .String()

// }

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string `json:"nombredesdejson"`
	Age      int
	Pass     bool
	Nickname string `json:"nick"`
}

// esto sirve para editar el string que sale cuando haces Println() de la variable
// func (s Student) String() string {
// 	return fmt.Sprintf("name: %s, age: %d, pass: %v, nick: %s", s.Name, s.Age, s.Pass, s.Nickname)
// }

func main() {
	jsonString := `{"nombredesdejson":"Lucas","Pass":false,"nick":"fune", "Age":27}`
	// creo una instancia de Student
	var student1 Student
	// a student1 le asigno los valores obtenidos del jsonString mediante json.Unmarshal()
	json.Unmarshal([]byte(jsonString), &student1)
	// si no le coloco el `json:"nombredesdejson"` y `json:"nick"` fmt.Println(student1) imprime asi: { 27 false }
	// y fmt.Println(student1.Name) por ejemplo no imprime nada
	fmt.Println(student1)

}
