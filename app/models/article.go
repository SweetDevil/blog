package models

type Article struct {
	CommonModel
	Id int
	Title string
	Ctime int
	Content string
	Etime int
	Thumb string
	Uid int
	Catid int
	Type int
}

func NewArticle() *Article {
	m := new(Article)
	return m
}