package controllers

import (
	"github.com/SweetDevil/blog/app/models"
	"github.com/revel/revel"
)

type ArticleAdmin struct {
	*revel.Controller
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
 */
func (c ArticleAdmin) Index() revel.Result {
	println(c.MethodName)
	c.RenderArgs["asd"] = 12
	return c.Render()
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
 */
func (c ArticleAdmin) Add() revel.Result {
	c.Request.ParseForm()
	c.Request.ParseMultipartForm(32 << 20)
	if c.Request.Method == "POST" {
		_,err:=models.ArticleModel.AddArticle(&c.Request.Form)
		//move uploaded file and write file path to db
		if err ==nil {
			models.ArticleModel.AddImage(&c.Request.MultipartForm.File)
		}
		//clear form
		c.Request.MultipartForm.RemoveAll()
	}
	c.RenderArgs["pageTitle"] = ""
	c.RenderArgs["category"] = models.CategoryModel.List()
	return c.RenderTemplate("Admin/Article/_form.html")
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
 */
func (c ArticleAdmin) Edit() revel.Result {
	println(c.MethodName)
	return c.Render()
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
 */
func (c ArticleAdmin) Del() revel.Result {
	println(c.MethodName)
	return c.Render()
}
