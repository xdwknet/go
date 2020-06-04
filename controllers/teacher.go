package controllers

import "github.com/astaxie/beego"

type TeacherController struct {
	beego.Controller
}

func (t *TeacherController) GetName() {
	t.Data["json"] = "hello world"
	t.ServeJSON()
}
