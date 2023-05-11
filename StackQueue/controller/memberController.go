package controller

import (
	"test/entity"
	"test/model"
)

func CreateFirst(data entity.Data){
	model.TambahAwal(data)
}

func CreateLast(data entity.Data){
	model.TambahAkhir(data)
}

func DeleteFirst(){
	model.HapusAwal()
}

func DeleteLast(){
	model.HapusAkhir()
}

func Show() []entity.Data {
	return model.Lihat()
}