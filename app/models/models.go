package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Connection gorm.DB

var ArticleModel *Article
var CategoryModel *Category

type CommonModel struct {
}

func init() {
	var err error
	Connection, err = gorm.Open("mysql", "root:723425@tcp(127.0.0.1:3306)/th?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println(err.Error())
	}
	Connection.DB().SetMaxIdleConns(10)
	Connection.DB().SetMaxOpenConns(100)
	Connection.SingularTable(true)

	ArticleModel = NewArticle()
	CategoryModel = NewCategory()
}
