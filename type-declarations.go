package main

import "fmt"

func main() {

	/**
	* type declaration adalah kemampuan untuk membuat tipe data baru dari tipe data yang sudah ada
	* type declaration biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada
	* dengan tujuan agar lebih mudah dimengerti
	 */

	type Married bool

	var burhanStatus Married = false
	var danaStatus Married = false
	var sintaStatus Married = true
	var jehanStatus Married = false

	fmt.Println(burhanStatus)
	fmt.Println(danaStatus)
	fmt.Println(sintaStatus)
	fmt.Println(jehanStatus)

}
