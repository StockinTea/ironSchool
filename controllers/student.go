package controllers

import (
  "ironSchool/models"

  "github.com/astaxie/beego"
)

type StudentController struct {
  beego.Controller
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