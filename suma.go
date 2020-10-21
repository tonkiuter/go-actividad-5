package main

import(
	"fmt"
)

//retornar el numero mas grande
func sumar(args ...int)int{
	total:= 0

	for _, v := range args{
		if v > total{
			total = v
		}
		
	}

	return total
}

func main(){

	a := []int{7, 4, 10, 90, 5}
	//fmt.Println(sumar(4, 5, 2))
	fmt.Println(sumar(a...)) // cada elemento del slice se envía como parámetros

}