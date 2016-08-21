package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout/layout.tpl"
	c.TplName = "main/index.tpl"
	c.Data["ActiveNav"] = "main"
}
