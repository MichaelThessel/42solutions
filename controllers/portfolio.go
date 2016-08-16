package controllers

import (
    "github.com/astaxie/beego"
    "42s/models/portfolio"
    "strconv"
)

type PortfolioController struct {
    beego.Controller
}

func (c *PortfolioController) Get() {
    c.Layout = "layout/layout.tpl"
    c.TplName = "portfolio/index.tpl"

    var p fortytwo.Portfolio
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

    if (id == 0) {
        id = 1
    }

    if (p.Get(id)) {
        c.Data["Portfolio"] = p
        c.Data["Next"] = p.Next()
        c.Data["Prev"] = p.Prev()
    } else {
        c.Abort("404")
    }
}
