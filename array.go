package main

import "fmt"

func main() {

	/**
	* pada go kita harus menentukan jumlah data yang bisa ditampung oleh array yang akan kita buat
	* daya tampung array ini adalah fix, artinya tidak bisa ditambah lagi
	* jika ingin menambah kapasitas datanya, harus membuat array baru
	 */
	var country [3]string
	country[0] = "indonesia"
	country[1] = "malaysia"
	country[2] = "singapura"

	/**
	* len digunakan untuk mengecek panjang array, walau array tersebut belum di input datanya,
	* len tetap akan mengembalikan ukuran si array, misal 10, berarti tetap 10,
	* walaupun tidak ada data yang di input
	 */
	fmt.Println(len(country))
	fmt.Println(country[0])
	fmt.Println(country[1])
	fmt.Println(country[2])
	fmt.Println()

	var price [3]int
	price[0] = 12000
	price[1] = 24000
	price[2] = 36000
	// price[3] = 72000 // akan menghasilkan error jika melebihi kapasitas daya tampung
	fmt.Println(price)
	fmt.Println()

	// membuat array secara langsung
	var decimal = [4]float32{
		3.4, 12.5, 1.7,
	}
	fmt.Println(decimal)
	fmt.Println()
}
