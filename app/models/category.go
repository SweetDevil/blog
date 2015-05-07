package models

import (
	"net/url"
)

type Category struct {
	CommonModel
	Id         int
	Name       string
	Pid        int
	Type       int
	Url        string
	Sort       int
	Background string
}

func NewCategory() *Category {
	m := new(Category)
	return m
}

func (m Category) List() (res []*Category) {
	Connection.Find(&res)
	return res
}

func (m Category) Add(form *url.Values) (id int, err error) {
	/*category := Category{
		Name:form.Get("category_name"),
	}*/
	return 1, nil
}
