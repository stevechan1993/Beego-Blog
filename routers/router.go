package routers

import (
	"github.com/astaxie/beego"
	"github.com/stevechan/Beego-Blog/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
