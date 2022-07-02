package main

import "fmt"

func main() {

	const ENV string = "DEVELOPMENT"

	// create multiple const
	const (
		DATABASE = "CRUD"
		USERNAME = "root"
		PASSWORD = ""
	)

	fmt.Println(ENV)
	fmt.Println(DATABASE)
	fmt.Println(USERNAME)
	fmt.Println(PASSWORD)

}
