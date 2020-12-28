package routers

import (
	"doc-lixiang/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})
	beego.Router("/documentParamList",&controllers.DocumentController{},"post,get:DocumentParamList")

}
