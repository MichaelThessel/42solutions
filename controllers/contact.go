package controllers

import (
	"github.com/MichaelThessel/42solutions/models/contact"
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) init() {
	c.Layout = "layout/layout.tpl"
	c.TplName = "contact/index.tpl"
	c.Data["ActiveNav"] = "contact"
}

func (c *ContactController) Get() {
	c.init()
}

func (c *ContactController) Post() {
	c.init()

	var m fortytwo.Message
	m.Name = c.Input().Get("name")
	m.Email = c.Input().Get("email")
	m.Telephone = c.Input().Get("telephone")
	m.Message = c.Input().Get("message")

	isValid, errors := m.Validate()
	c.Data["Message"] = m
	c.Data["IsValid"] = isValid
	c.Data["Errors"] = errors

	if isValid {
		m.Send()
	}
}
