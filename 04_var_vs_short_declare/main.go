package main

import "fmt"

func main() {
	var city string
	city = "Delhi"
	// city is a variable of type string and it holds the value "Delhi"

	var channels = "bengaluru" //infering the type of variable based on the value assigned to it which is string in this case

	house_num := 66                //short declaration operator which is used to declare and initialize a variable in one line, it infers the type of variable based on the value assigned to it which is int in this case
	faltmate, age := "bodhiji", 23 //short declaration operator can be used to declare and initialize multiple variables in one line, it infers the type of variable based on the value assigned
	//to it which is string and int in this case
	fmt.Println("Hi my name is aditya and I'm from", city, "but i live in", channels, "\n my house is: ", house_num, "/\b my house mate is: ", faltmate, "\n his age is: ", age)

}
