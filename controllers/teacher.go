package controllers

import (
  "ironSchool/models"

  "github.com/astaxie/beego"
)

type TeacherController struct {
  beego.Controller
}

func (this *TeacherController) GetAllTeachers() {
  teacheres := models.GetAllTeachers()
  this.Data["json"] = teacheres
  this.ServeJSON()
}

func (this *TeacherController) PostNewTeacher() {
  tName := this.GetString("TeacherName")
  tAge, _ := this.GetInt("TeacherAge")
  tTel := this.GetString("TeacherTel")
  newTeacherId := models.AddNewTeacher(models.Teacher{"", tName, tAge, tTel})
  this.Data["json"] = map[string]string{"data": newTeacherId}
  this.ServeJSON()
}

func (this *TeacherController) DelTeacherByTID() {
  tid := this.GetString(":teacherId")
  models.DeleteTeacher(tid)
  this.Data["json"] = map[string]string{"status": "success", "data": tid}
  this.ServeJSON()
}
