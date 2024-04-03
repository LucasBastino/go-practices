package main

import (
	"fmt"
)

type Usuario struct {
	nombre  string
	edad    int
	amigos  []Usuario
	posteos map[string]string
	activo  bool
}

type Admin struct {
	usuarios []Usuario
}

func main() {
	admin := Admin{}
	lucas := admin.crearYGuardar(Usuario{"Lucas", 27, []Usuario{}, map[string]string{}, true})
	jorge := admin.crearYGuardar(Usuario{"Jorge", 25, []Usuario{}, map[string]string{}, true})
	luana := admin.crearYGuardar(Usuario{"Luana", 24, []Usuario{}, map[string]string{}, true})
	lucas.agregarAmigo(jorge)
	lucas.agregarAmigo(luana)
	lucas.publicarPost("Titulo 1", "este es el primer posteo")
	lucas.publicarPost("Titulo 2", "este es el segundo posteo")
	lucas.imprimirInfo()
	lucas.borrarPost("Titulo 1")
	lucas.imprimirInfo()
	admin.imprimirUsuarios()
	admin.promedioEdad()
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
	for titulo, post := range u.posteos {
		fmt.Printf("%s:\n%s\n", titulo, post)
	}
}

// func (u *Usuario) modificarAtributo(atributo string, valor any) {
// 	fmt.Println(u.atributo)
// }

func (u *Usuario) modificarNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) activarDesactivar() {
	if u.activo {
		u.activo = false
	} else {
		u.activo = true
	}
}

func (u *Usuario) agregarAmigo(nuevoAmigo Usuario) {
	u.amigos = append(u.amigos, nuevoAmigo)
}

// func (u *Usuario) eliminarAmigo(amigo Usuario) {

// 	sliceAmigos := []string{}
// 	for _, value := range u.amigos {
// 		sliceAmigos = append(sliceAmigos, value.nombre)
// 	}

// 	index := sliceMethods.FindIndex(sliceAmigos, amigo.nombre)
// 	u.amigos = append(u.amigos[:index], u.amigos[index+1:]...)
// }

func (u *Usuario) eliminarAmigo(amigo Usuario) {
	for i, value := range u.amigos {
		if value.nombre == amigo.nombre {
			u.amigos = append(u.amigos[:i], u.amigos[i+1:]...)
			fmt.Printf("Usuario %v eliminado de tu lista de amigos\n", value.nombre)
			return
		}
	}
	fmt.Println("No se ha encontrado el usuario a eliminar")
}

func (u *Usuario) publicarPost(titulo string, post string) {
	u.posteos[titulo] = post
}

func (u *Usuario) borrarPost(titulo string) {
	delete(u.posteos, titulo)
}

func (a *Admin) crearYGuardar(nuevoUsuario Usuario) Usuario {
	a.usuarios = append(a.usuarios, nuevoUsuario)
	return nuevoUsuario
}

func (a Admin) imprimirUsuarios() {
	for _, usuario := range a.usuarios {
		fmt.Println(usuario.nombre)
	}
}

func (a Admin) promedioEdad() {
	sumaTotal := 0
	for _, usuario := range a.usuarios {
		sumaTotal += usuario.edad
	}
	fmt.Println("Promedio de edad:", sumaTotal/len(a.usuarios))
}
