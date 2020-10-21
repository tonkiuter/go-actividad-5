package main

import "fmt"

//Crea un programa que pida dos enteros (a y b) y llame a la función intercambia(&a, &b), 
//la cual intercambia el valor de a y b; usando punteros.
func intercambia(a *int, b *int) {
	aux := *b
	*b = *a
	*a = aux
}

func main(){
	var a,b int
	
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	intercambia(&a,&b)

	fmt.Println(a)
	fmt.Println(b)

}