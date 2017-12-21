package routers

import (
	"Board/controllers"

	"github.com/astaxie/beego"
)

func init() {
beego.Router("/", &controllers.StudentController{},"*:GetAll")
beego.Router("/:studentId", &controllers.StudentController{},"*:Get")
}