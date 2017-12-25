package controllers

import (
  "encoding/json"
  // "fmt"
  "github.com/astaxie/beego"
)

type DemoController struct {
  beego.Controller
}

type Demo struct {
  Status  string `json:stauts`
  IsAdult bool   `json:isAdult`
}

func (this *DemoController) PutRequestBodyTest() {
  newD := Demo{}
  json.Unmarshal(this.Ctx.Input.RequestBody, &newD)
  this.Data["json"] = newD
  this.ServeJSON()
}
