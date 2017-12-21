package controllers

import (
  "ironSchool/models"

  "github.com/astaxie/beego"
)

type ClassController struct {
  beego.Controller
}

func (this *ClassController) GetById() {
  classId := this.Ctx.Input.Param(":classId")
  if classId != "" {
    class, err := models.GetClassById(classId)
    if err != nil {
      this.Data["json"] = err.Error()
    } else {
      this.Data["json"] = class
    }
  }
  this.ServeJSON()
}

func (this *ClassController) GetByName() {
  className := this.Ctx.Input.Param(":className")
  if className != "" {
    class, err := models.GetClassByName(className)
    errorResp := make(map[string]string)
    if err != nil {
      errorResp["status"] = err.Error()
      this.Data["json"] = errorResp
    } else {
      this.Data["json"] = class
    }
  }
  this.ServeJSON()
}

func (this *ClassController) GetAllClass() {
  classes := models.GetAllClasses()
  this.Data["json"] = classes
  this.ServeJSON()
}