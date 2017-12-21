package routers

import (
  "ironSchool/controllers"

  "github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.StudentController{},"*:GetAll")
  beego.Router("/:studentId", &controllers.StudentController{},"*:Get")
}