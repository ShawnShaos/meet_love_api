package controllers

import (
	"github.com/astaxie/beego"
	"jie_ling_api/models"
	"jie_ling_api/untils"
)

type UsersController struct {
	beego.Controller
}

func (u *UsersController) GetUser() {
	id, _ := u.GetInt("id")
	user,err := models.GetUserById(id)
	if err {
		u.Data["json"] = untils.GenTip(-1, "fail", "")
	}else {
		u.Data["json"] = untils.GenTip(1, "ok", user)
	}
	u.ServeJSON()
}
