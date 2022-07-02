package main

import "fmt"

func main()  {
	
	/**
	* saat membuat variable, maka kita wajib mendefinisikan tipe datanya
	* namun, jika kita langsung menginisialisasikan datanya pada var tersebut, maka kita tidak wajib menyebutkan tipe datanya
	*/

	var country string;
	var city string;
	var peopleCount int;

	country = "indonesia";
	city = "jakarta";
	peopleCount = 20000000;
	fmt.Println(country);
	fmt.Println(city);
	fmt.Println(peopleCount);
	fmt.Println("======================");

	country = "malaysia";
	city = "kuala lumpur";
	peopleCount = 35700000;
	fmt.Println(country);
	fmt.Println(city);
	fmt.Println(peopleCount);
	fmt.Println("======================");

	/**
	* initiate variabel langsung, tanpa harus mendefinisikan tipe datanya terlebih dahulu
	* syaratnya, anda harus langsung memberikan nilai pada variable tersebut
	*/
	var text = "Hello Country";
	fmt.Println(text);
	fmt.Println("======================");

	/**
	* kita juga bisa membuat variable tanpa menggunakan kata kunci var,
	* caranya adalah dengan menggunakan simbol :=
	* namun cara ini hanya saat pendefinisian variable di awal saja
	*/
	continent := "Asia Tenggara";
	fmt.Println(continent);

	continent = "Eropa";
	fmt.Println(continent);
	fmt.Println("======================");

	/**
	* di golang kita juga bisa mendeklarasikan multiple variable
	*/
	var(
		country1 = "Myanmar";
		country2 = "Laos";
		country3 = "Vietnam";
		country4 = "Thailand";
	)
	fmt.Println(country1);
	fmt.Println(country2);
	fmt.Println(country3);
	fmt.Println(country4);
	fmt.Println("======================");

}