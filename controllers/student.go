package controllers

import (
  "ironSchool/models"

  "github.com/astaxie/beego"
)

type StudentController struct {
  beego.Controller
}

func (this *StudentController) Get() {
  StudentId := this.Ctx.Input.Param(":studentId")
  if StudentId != "" {
    st, err := models.GetStudent(StudentId)
    if err != nil {
      this.Data["json"] = err.Error()
    } else {
      this.Data["json"] = st
    }
  }
  this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
  this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
  this.ServeJSON()
}

func (this *StudentController) GetAll() {
  this.Ctx.ResponseWriter.Header().Set("ServerMan", "StockinTea")
  sts := models.GetAllStudents()
  this.Data["json"] = sts
  this.ServeJSON()
}

func (this *StudentController) PostNewStudent() {
  var newStudent models.Student
  this.ParseForm(&newStudent)
  newStudentId := models.AddNewStudent(newStudent)
  this.Data["json"] = map[string]string{"StudentId": newStudentId}
  this.ServeJSON()
}