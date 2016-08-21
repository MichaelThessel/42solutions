package routers

import (
	"42s/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/portfolio/?:id", &controllers.PortfolioController{})
}
