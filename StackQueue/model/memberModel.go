package model

import (
	"test/database"
	"test/entity"
)

func TambahAwal(data entity.Data) {
	temp := &database.DataMember

	new := entity.Member{
		Data: entity.Data{
			Nama:  data.Nama,
			Point: data.Point,
		},
	}

	if temp.Next == nil {
		temp.Next = &new
	} else {
		new.Next = temp.Next
		temp.Next = &new
	}
}

func TambahAkhir(data entity.Data) {
	temp := &database.DataMember
	new := entity.Member{
		Data: data,
	}

	if temp.Next == nil {
		temp.Next = &new
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &new
	}
}


func HapusAwal() {
	temp := &database.DataMember

	if temp.Next == nil {
		return
	} else {
		temp.Next = temp.Next.Next
	}
}

func HapusAkhir() {
	temp := &database.DataMember

	if temp.Next == nil {
		return
	} else {
		for temp.Next != nil {
			if temp.Next.Next == nil {
				temp.Next = nil
				break
			}
			temp = temp.Next
		}
	}
}

func Lihat() []entity.Data {
	temp := database.DataMember.Next
	data := []entity.Data{}

	for temp != nil {
		data = append(data, temp.Data)
		temp = temp.Next
	}

	return data
}