package model

import (
	"stack/database"
	"stack/entity"
)

func Push(data entity.Data) {
	temp := &database.DataMember
	new := entity.Member{Data: data}

	if temp.Next == nil {
		temp.Next = &new
	} else {
		new.Next = temp.Next
		temp.Next = &new
	}
}

func Pop() bool {
	temp := &database.DataMember

	if temp.Next == nil {
		return false
	}

	prev := temp
	cur := temp.Next

	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}

	prev.Next = nil
	return true
}

func Find() []entity.Data {
	temp := &database.DataMember
	var data []entity.Data
	cur := temp.Next

	for cur != nil {
		data = append(data, cur.Data)
		cur = cur.Next
	}

	return data
}