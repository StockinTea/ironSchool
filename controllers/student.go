package controllers

import (
  "Board/models"
  "encoding/json"

  "github.com/astaxie/beego"
)

type StudentController struct {
  beego.Controller
}


func (s *StudentController) Post() {
  var st models.Student
  json.Unmarshal(s.Ctx.Input.RequestBody, &st)
  StudentId := models.AddStudent(st)
  s.Data["json"] = map[string]string{"StudentId": StudentId}
  s.ServeJSON()
}

func (s *StudentController) Get() {
  StudentId := s.Ctx.Input.Param(":studentId")
  if StudentId != "" {
    st, err := models.GetStudent(StudentId)
    if err != nil {
      s.Data["json"] = err.Error()
    } else {
      s.Data["json"] = st
    }
  }
  s.ServeJSON()
}

func (s *StudentController) GetAll() {
  sts := models.GetAllStudents()
  s.Data["json"] = sts
  s.ServeJSON()
}