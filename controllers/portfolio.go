package controllers

import (
	"github.com/MichaelThessel/42solutions/models/portfolio"
	"github.com/astaxie/beego"
	"strconv"
)

type PortfolioController struct {
	beego.Controller
}

func (c *PortfolioController) Get() {
	c.Layout = "layout/layout.tpl"
	c.TplName = "portfolio/index.tpl"
	c.Data["ActiveNav"] = "portfolio"

	var p fortytwo.Portfolio
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	if id == 0 {
		id = 1
	}

	if p.Get(id) {
		c.Data["Portfolio"] = p
		c.Data["Next"] = p.Next()
		c.Data["Prev"] = p.Prev()
	} else {
		c.Abort("404")
	}
}
