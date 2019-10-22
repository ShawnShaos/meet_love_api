// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"jie_ling_api/controllers"
	"jie_ling_api/filters"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	beego.Router("/idcardauth",&controllers.UserAuthController{},"post:IdCardAuth")
	beego.Router("/wxlogin", &controllers.UserController{}, "post:WxLogin")   //登录
	beego.Router("/wxUserInfo", &controllers.UserController{}, "post:WxUserInfo")   //登录并获取用户加密信息（第一次使用）
	//beego.Router("/login", &controllers.UserController{}, "post:Login")   //登录
	//beego.Router("/registed", &controllers.UserController{}, "post:Registed")   //注册
	beego.Router("/verify", &controllers.VerifyController{})   //验证码
	beego.Router("/users/GetUserPic",&controllers.UserController{},"get:GetUserPic")
	beego.Router("/GetUserPic", &controllers.UserController{},"get:GetUserPic")   //验证码

	user := beego.NewNamespace("/user").
		Filter("before",filters.Auth()).
		Router("/index", &controllers.UserController{}, "get:Index")      //首页

	beego.AddNamespace(user)
}
