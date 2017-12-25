package routers

import (
  "ironSchool/controllers"

  "github.com/astaxie/beego"
)

func init() {
  namespace := beego.NewNamespace("/ironSchool",
    beego.NSNamespace("/student",
      beego.NSRouter("/", &controllers.StudentController{}, "get:GetAll"),
      beego.NSRouter("/", &controllers.StudentController{}, "post:PostNewStudent"),
      beego.NSRouter("/:studentId", &controllers.StudentController{}, "*:Get"),
    ),
    beego.NSNamespace("/class",
      beego.NSRouter("/", &controllers.ClassController{}, "get:GetAllClass"),
      beego.NSRouter("/", &controllers.ClassController{}, "post:PostNewClass"),
      beego.NSRouter("/:classId", &controllers.ClassController{}, "get:GetById"),
      beego.NSNamespace("/name",
        beego.NSRouter("/:className", &controllers.ClassController{}, "*:GetByName"),
      ),
    ),
    beego.NSNamespace("/teacher",
      beego.NSRouter("/", &controllers.TeacherController{}, "get:GetAllTeachers"),
      beego.NSRouter("/", &controllers.TeacherController{}, "post:PostNewTeacher"),
      beego.NSRouter("/:teacherId", &controllers.TeacherController{}, "delete:DelTeacherByTID"),
    ),
    beego.NSNamespace("/test",
      beego.NSRouter("/", &controllers.DemoController{}, "put:PutRequestBodyTest"),
    ),
  )
  beego.AddNamespace(namespace)
}
