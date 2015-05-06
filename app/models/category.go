package models

type Category struct {
	CommonModel
	Id int
	Name int
	Pid int
	Type int
	Url string
	Sort int
	Background string
}

func NewCategory() *Category {
	m := new(Category)
	return m
}