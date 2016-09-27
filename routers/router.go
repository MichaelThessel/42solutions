package routers

import (
	"github.com/MichaelThessel/42solutions/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/portfolio/?:id", &controllers.PortfolioController{})
}
