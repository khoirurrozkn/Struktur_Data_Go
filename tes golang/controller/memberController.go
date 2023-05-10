package controller

import (
	"stack/entity"
	"stack/model"
)

func Create(data entity.Data) {

	model.Push(data)
}

func Delete() {
	model.Pop()
}

func Show() []entity.Data {
	return model.Find()
}
