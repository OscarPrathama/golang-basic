package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)

	/**
	* hati2 ketika melakukan konversi, karena misal terdapat case dimana kita akan mengkonversi
	* nilai 129 ke int8, yang dimana nilai max int8 adalah 127, maka akan terjadi interger overflow
	* dimana nilai int8 akan kembali dari awal ke selanjutnya yakni -128, dst...
	* coba rubah nilai32 ke 127 / 128 / 129, dan perhatikan pada variable nilai8
	 */
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println("=================")

	/**
	* Saat kita mengambil karakter dari suatu String, tipe defaultnya adalah byte,
	* maka anda harus mengkonversikannya ke string untuk mendapatkan karakternya.
	* untuk membuktikannya, silahkan arahkan kursor ke var getString, tipe datanya pasti byte
	 */
	var country = "indonesia"
	var getString = country[2]
	fmt.Println(getString)
	var getRealString string = string(getString)
	fmt.Println(getRealString)

}
