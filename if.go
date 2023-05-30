package main

import "fmt"

func main() {
	nama := "alung"

	if nama == "alung" {
		fmt.Println("anjay alung")
	}else{
		fmt.Println("Bukan alung")
	}
	sayHello()
	result := terimaNomor(2,4)
	fmt.Println(result)
}

func sayHello(){
	fmt.Println("halo ini function lulung")
}

func terimaNomor(a,b int)int{
	return a + b
}