package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var Connection gorm.DB

var ArticleModel *Article
var CategoryModel *Category

type CommonModel struct {
	
}

func init() {
	var err error
	Connection,err = gorm.Open("mysql", "root:723425@tcp(127.0.0.1)/th?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		println(err.Error())
	}
	ArticleModel = NewArticle()
	CategoryModel = NewCategory()
}