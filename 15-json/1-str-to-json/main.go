// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // los atributos tienen que ser exportables para que json los pueda usar, por eso empiezan con mayuscula
// type Student struct {
// 	Name     string `json:"nombre"`
// 	Age      int    `json:"Age,omitempty"`
// 	Pass     bool
// 	Nickname string `json:"nick,omitempty"`
// }

// func main() {
// 	student1 := Student{
// 		Name:     "Lucas",
// 		Age:      20,
// 		Pass:     true,
// 		Nickname: "",
// 	}
// 	studentInJson, err := json.Marshal(student1)
// 	if err != nil {
// 		fmt.Println("error converting str to json")
// 		panic(err.Error())
// 	}
// 	fmt.Println(studentInJson)
// 	studentInString := string(studentInJson)
// 	fmt.Println(studentInString)
// }

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string `json:"nombre desde json"`
	Age      int    `json:"edad,omitempty"`
	Pass     bool
	Nickname string
}

func main() {
	student1 := Student{
		Name:     "Lucas",
		Age:      0, // 0 en int es empty
		Pass:     false,
		Nickname: "fune",
	}

	studentInJson, err := json.Marshal(student1)
	if err != nil {
		fmt.Println("error marshaling student")
		panic(err.Error())
	}
	studentInString := string(studentInJson)
	fmt.Println(studentInJson)
	fmt.Println(studentInString)

}
