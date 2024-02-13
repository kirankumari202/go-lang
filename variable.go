package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var username string = "kiran"
	fmt.Println(username)
	fmt.Printf("type of the variable of user: %T \n", username)

	var isvalid bool = true
	fmt.Println(isvalid)
	fmt.Printf("type of the variable of user: %T \n", isvalid)

	var num int = 10000
	fmt.Println(num)
	fmt.Printf("type of the variable of user: %T \n", num)

	var largenum float32 = 10.23445
	fmt.Println(largenum)
	fmt.Printf("type of the variable of user: %T \n", largenum)

	//default value of int
	var dnum int
	fmt.Println(dnum)

	var number int = 10
	fmt.Println(number)

	//implicit declaration
	var name = "kiranio"
	fmt.Println(name)

	//no var declaration
	Newnumber := 40
	fmt.Println(Newnumber)

	//take variable from user
	welcome := "welcome to the take variable from user"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza")

	//what ever reader read the data all are going to sore here by using comma ok, err ok

}
