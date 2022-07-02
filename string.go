package main

import "fmt"

func main() {
	fmt.Println("Hello String")
	fmt.Println("Hello World")
	fmt.Println("Hello Earth")

	fmt.Println(len("Hello String")) // output adalah 12, yang merupakan panjang dari string
	fmt.Println("Hello String"[0])   // outputnya adalah 72, yang merupakan representasi byte huruf H
	fmt.Println("Hello String"[1])   // outputnya adalah 101, yang merupakan representasi byte huruf e
}
