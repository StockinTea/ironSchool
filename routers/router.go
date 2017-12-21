package routers

import (
  "ironSchool/controllers"

  "github.com/astaxie/beego"
)

func init() {
  namespace := beego.NewNamespace("/ironSchool",
    beego.NSNamespace("/student",
      beego.NSRouter("/", &controllers.StudentController{},"*:GetAll"),
      beego.NSRouter("/:studentId", &controllers.StudentController{},"*:Get"),
    ),
    beego.NSNamespace("/class",
      beego.NSRouter("/", &controllers.ClassController{},"*:GetAllClass"),
      beego.NSRouter("/:classId", &controllers.ClassController{},"*:GetById"),
      beego.NSNamespace("/name",
        beego.NSRouter("/:className", &controllers.ClassController{},"*:GetByName"),
      ),
    ),
  )
  beego.AddNamespace(namespace)
}