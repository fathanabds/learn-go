package main

import "fmt"

type Address struct {
	city, province, country string
}

func changeCountryToIndonesia(address *Address) {
	address.country = "indonesia"
}

func main() {
	address1 := Address{"bekasi", "jawa barat", "indonesia"}
	address2 := address1  // pass by value
	address3 := &address1 // pass by reference

	address2.city = "bandung"
	address3.city = "karawang"

	*address3 = Address{"malang", "jawa timur", "indonesia"} // memindahkan semua pointer yang mengacu pada address1 menjadi ke data yang baru yaitu address3
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	fmt.Println("=====")

	alamat1 := new(Address)
	fmt.Println(alamat1)

	alamat2 := alamat1

	alamat2.country = "indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	fmt.Println("=====")

	alamat := Address{city: "subang", province: "jawa barat", country: ""}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
