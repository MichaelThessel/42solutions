package controllers

import (
    "github.com/astaxie/beego"
)

type PortfolioController struct {
    beego.Controller
}

func (c *PortfolioController) Get() {
    c.Layout = "layout/layout.tpl"
    c.TplName = "portfolio/index.tpl"
}
