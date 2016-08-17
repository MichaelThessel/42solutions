package controllers

import (
    "github.com/astaxie/beego"
)

type ContactController struct {
    beego.Controller
}

func (c *ContactController) Get() {
    c.Layout = "layout/layout.tpl"
    c.TplName = "contact/index.tpl"
    c.Data["ActiveNav"] = "contact";
}
