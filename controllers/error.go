package controllers

import (
  "github.com/astaxie/beego"
)

type ErrorController struct {
  beego.Controller
}

func (this *ErrorController) Error404() {
  this.Data["json"] = map[string]string{"msg": "page not found", "status": "error404"}
  this.ServeJSON()
}

func (this *ErrorController) Error500() {
  this.Data["json"] = map[string]string{"msg": "internal server error", "status": "error500"}
  this.ServeJSON()
}