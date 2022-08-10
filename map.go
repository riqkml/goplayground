package main

import "fmt"

func main() {
	var listMakanan = map[int]string{10: "nasi", 20: "roti", 30:"mie"}

	for idx, makanan := range listMakanan {
		fmt.Println("kode makanan",idx,"dengan nama", makanan)
	}
	delete(listMakanan, 10)

	for idx, makanan := range listMakanan {
		fmt.Println("kode update makanan",idx,"dengan nama", makanan)
	}
}