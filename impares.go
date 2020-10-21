package main

import "fmt"

//Crea un generador de número impares tomando de ejemplo generadorPares().
func generadorImpares() func() uint{
	i := uint(1)

	return func() uint{
		var impar = i
		i += 2
		return impar
	}
}


func main(){
	nextimpar := generadorImpares()

	fmt.Println(nextimpar())
	fmt.Println(nextimpar())
	fmt.Println(nextimpar())
	fmt.Println(nextimpar())

}