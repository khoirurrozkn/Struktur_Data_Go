package view

import (
	"fmt"
	"stack/controller"
	"stack/entity"
)

func Hero() {
	var nama string
	var point int

	fmt.Println("Masukan nama : ")
	fmt.Scan(&nama)
	fmt.Println("Masukan point : ")
	fmt.Scan(&point)

	data := entity.Data{
		Nama:  nama,
		Point: point,
	}
	controller.Create(data)

	fmt.Println("Masukan nama : ")
	fmt.Scan(&nama)
	fmt.Println("Masukan point : ")
	fmt.Scan(&point)

	data2 := entity.Data{
		Nama:  nama,
		Point: point,
	}
	controller.Create(data2)

	fmt.Println("Masukan nama : ")
	fmt.Scan(&nama)
	fmt.Println("Masukan point : ")
	fmt.Scan(&point)

	data3 := entity.Data{
		Nama:  nama,
		Point: point,
	}
	controller.Create(data3)

	controller.Delete()

	ea := controller.Show()
	for i,v := range ea {
		fmt.Println("urutan ke : ", i)
		fmt.Println("nama ", v.Nama)
		fmt.Println("point", v.Point)
	}

}
