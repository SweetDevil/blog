package controllers

import (
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
func (c ArticleAdmin) Index() revel.Result{
	println(c.MethodName)
	c.RenderArgs["asd"] = 12
	return c.Render()
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
*/
func (c ArticleAdmin) Add() revel.Result{
	
	if c.Request.Method == "POST" {
		
	}
	c.RenderArgs["abc"] = 12
	return c.RenderTemplate("Admin/Article/_form.html")
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
*/
func (c ArticleAdmin) Edit() revel.Result{
	println(c.MethodName)
	return c.Render()
}

/**
* add an article
* @author yuxing951@hotmail.com
* @date 2015-04-28
*/
func (c ArticleAdmin) Del() revel.Result{
	println(c.MethodName)
	return c.Render()
}