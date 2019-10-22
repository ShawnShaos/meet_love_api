package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/dchest/captcha"
)

type VerifyController struct {
	beego.Controller
}

//func (v *VerifyController) Get() {
//	d:= struct {
//		CaptchaId string
//	}{
//		captcha.New(),
//	}
//
//	v.Data["captchaid"] = d
//	v.ServeJSON()
//}