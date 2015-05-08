package models

import (
	"github.com/SweetDevil/blog/app/tplfunc"
	"mime/multipart"
	"net/url"
)

type Article struct {
	CommonModel
	Id      int
	Title   string
	Ctime   int
	Content string
	Etime   int
	Thumb   string
	Uid     int
	Catid   int
	Type    int
}

func NewArticle() *Article {
	m := new(Article)
	return m
}

//add an article
//@date 2015-05-06
//@author yuxing951@hotmail.com
func (m Article) AddArticle(form *url.Values) (id int, err error) {
	article := Article{
		Title:   form.Get("article_title"),
		Ctime:   tplfunc.TimeStamp(),
		Content: form.Get("article_content"),
		Etime:   tplfunc.TimeStamp(),
		Thumb:   form.Get("article_thumb"),
		Uid:     tplfunc.CurrentUid(),
		Catid:   tplfunc.Intval(form.Get("article_catid")),
		Type:    tplfunc.Intval(form.Get("article_type")),
	}
	r := Connection.Create(&article)

	return 1, r.Error
}

//add article image
//@date 2015-05-07
//@author yuxing951@hotmail.com
func (m Article) AddImage(f *map[string][]*multipart.FileHeader) (err error) {
	uploadPath := tplfunc.ImageUploadPath()
	err = tplfunc.TryMkdir(uploadPath)
	files := *f
	thumb := files["article_thumb"]
	for _, v := range thumb {
		file, _ := v.Open()
		tplfunc.MoveUploadFile(file, uploadPath+"/"+"xxx.jpg")
	}
	return nil
}
