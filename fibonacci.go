package main

import "fmt"

//Serie de Fibonacci
func fibo(n int) int{
	if n == 0 || n == 1{
		return n
	
}else{
		return fibo(n-1) + fibo(n-2)
	}
	
}


func main(){
	var n int

	fmt.Scanln(&n)

	for i:= 0; i < n; i++{
		fmt.Print(fibo(i),",")
	} 

}