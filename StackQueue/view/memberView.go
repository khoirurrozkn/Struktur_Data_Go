package view

import (
	"fmt"
	"test/entity"
	"test/controller"
)

func Hero(){

	var nama string
	var point int
	var pilihan int
	var data entity.Data
	for{
		fmt.Println("Masukan pilihan")
		fmt.Println("1. tambah awal")
		fmt.Println("2. tambah akhir")
		fmt.Println("3. hapus awal")
		fmt.Println("4. hapus akhir")
		fmt.Println("5. lihat semua data")
		fmt.Scan(&pilihan)

		if pilihan == 1 || pilihan == 2 {
			fmt.Println("Masukan nama anda : ( string )")
			fmt.Scan(&nama)
			fmt.Println("Masukan poin anda : ( int )")
			fmt.Scan(&point)

			data = entity.Data{
				Nama: nama,
				Point: point,
			}
		}

		if pilihan == 1 {
			controller.CreateFirst(data)
		}else if pilihan == 2 {
			controller.CreateLast(data)
		}else if pilihan == 3 {
			controller.DeleteFirst()
		}else if pilihan == 4 {
			controller.DeleteLast()
		}else {
			ea := controller.Show()

			for i,val := range ea {
				fmt.Println("urutan ke : ", i)
				fmt.Println("nama      : ", val.Nama)
				fmt.Println("poin      : ", val.Point, "\n")
			}
		}
	}
}