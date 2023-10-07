package main

import "fmt"

func makeCar(car map[string]string) string {
	return "Mobil " + car["name"] + " berwarna " + car["color"]
}

func result(message string) {
	fmt.Println(message)
}

func main() {
	var car = make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"

	message := makeCar(car)
	result(message)

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// output => Mobil BMW berwarna Black

}
