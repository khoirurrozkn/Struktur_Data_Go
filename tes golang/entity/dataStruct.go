package entity

type Data struct {
	Nama  string
	Point int
}

type Member struct {
	Data Data
	Next *Member
}
