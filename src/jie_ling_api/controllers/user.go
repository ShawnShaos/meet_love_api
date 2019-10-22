package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"io/ioutil"
	"jie_ling_api/filters"
	"jie_ling_api/models"
	"jie_ling_api/untils"
	"net/http"
	"jie_ling_api/untils/wxbizdatacrypt"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

const (
	APPID  = "wx50e6fd332b3aebf6"
	SECRET = "7aefac47f5ddeac3c902299d6384ed2c"
)

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
//func (u *UserController) Login() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	if models.Login(username, password) {
//		ret := filters.CreatToken(username,123)
//		if ret["code"] == 1 {
//			u.Ctx.Output.Header("Authorization", ret["data"].(string)) //设置登录之后返回的token
//		} else {
//			u.Ctx.Output.Header("Authorization", "") //设置登录之后返回的token
//		}
//		u.Data["json"] = untils.GenTip(1, "login success", username)
//	} else {
//		u.Data["json"] = untils.GenTip(-1, "user not exist", "")
//	}
//	u.ServeJSON()
//}

/**
注册
 */
//func (u *UserController) Registed() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	repassword := u.GetString("repassword")
//
//	if password != repassword {
//		u.Data["json"] = untils.GenTip(-1, "第一次和第二次输入的密码不一致", "")
//	}
//
//	userinfo := new(models.UserInfo)
//	userinfo.Name = username
//	userinfo.Password = password
//
//	ret, err := models.AddUserInfo(*userinfo)
//	if err != nil {
//		u.Data["json"] = untils.GenTip(-1, err.Error(), "")
//	} else { //注册成功，返回用户token
//		token := filters.CreatToken(username,123)
//		u.Ctx.Output.Header("Authorization", token["data"].(string))
//		u.Data["json"] = untils.GenTip(-1, "", int64(ret))
//	}
//	u.ServeJSON()
//}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

func (u *UserController) Index() {
	u.Data["json"] = "Index success"
	u.ServeJSON()
}

/**
获取验证码
 */
func (u *UserController) Captcha() {

	image := captcha.NewImage([]byte("s133"), 100, 100)
	u.Data["image"] = image
	u.ServeJSON()
}

/**
微信小程序登录
 */
func (u *UserController) WxLogin() {
	code := u.GetString("code")

	resp, _ := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + code + "&grant_type=authorization_code")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var bodyMap map[string]string
    json.Unmarshal(body,&bodyMap)

	userAuths := models.UserAuths{Identifier: bodyMap["openid"]}  //查询是否有该用户
	userInfo, err := models.GetUserByOpenId(userAuths)
	fmt.Println(userInfo)
	fmt.Println(err)
	if err !=nil{ //用户不存在，
		u.Data["json"] = untils.GenTip(-1, "用户不存在", "")
	}else {  //用户存在，返回jwt验证

		token := filters.CreatToken(filters.UserJwtBase{userInfo["nickname"].(string),userInfo["uid"].(int)})  //生成token
		u.Ctx.Output.Header("Authorization", token["data"].(string))
		u.Data["json"] = untils.GenTip(1, "登录成功", "")
	}
	u.ServeJSON()
}

/**
第一次或识别不了用户时重新获取用户加密信息
 */
func (u *UserController) WxUserInfo() {
	code := u.GetString("code")
	encryptedData := u.GetString("encryptedData")
	iv := u.GetString("iv")
	resp, _ := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + code + "&grant_type=authorization_code")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var userAuth map[string]interface{}

	json.Unmarshal(body, &userAuth) //字节数组转map

	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: APPID, SessionKey: userAuth["session_key"].(string)}
	result, _ := pc.Decrypt(encryptedData, iv, false) //获取到用户的加密信息 map[avatarUrl:https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eoJc0NdcPAOf5h2ldehOgibCHN187Zfmms9bQXpsIbLwAjibXyeK4Mo5lLsXKK89ibaMId4w2XWamTKA/132 nickName:LY language:zh_CN province:Sichuan country:China watermark:map[timestamp:1.558173247e+09 appid:wx50e6fd332b3aebf6] openId:oTGv-0NH7yq8jCZuIZYo6070tfJQ gender:1 city:Chengdu]
	fmt.Println(result)
	resultJson := result.(map[string]interface{})
	fmt.Println(resultJson["openId"].(string))
	userAuths := models.UserAuths{Identifier: resultJson["openId"].(string)}
	res, err := models.GetUserByOpenId(userAuths)
	fmt.Println(res)
	fmt.Println(err)
	if err == nil { //该用户已存在，直接登录
		fmt.Println("用户存在",res)
		u.Data["json"] = res
	} else { //不存在，新用户

		var userBase models.UserBase
		userBase.Nickname = resultJson["nickName"].(string) //昵称
		userBase.Avatar = resultJson["avatarUrl"].(string)  //头像
		userBase.Country = resultJson["country"].(string)   //国家
		userBase.City = resultJson["city"].(string)         //城市
		userBase.Province = resultJson["province"].(string) //省份
		userBase.CreateTime = time.Now() //创建时间
		userBase.LoginTime = time.Now() //登录时间
		userBase.LastLoginTime = time.Now()  //最后登录时间

		idUserBase:= models.InsertUserBase( userBase)
		fmt.Println("userBase的Id为：",idUserBase)

		var userAuths models.UserAuths
		userBase.Id = int(idUserBase)
		userAuths.UserBase = &userBase
		userAuths.IdentityType = "wx" //微信注册标识别
		userAuths.Identifier = resultJson["openId"].(string)  //openId

		idUserAuth := models.InsertUserAuths( userAuths)
		fmt.Println("userAuths的Id为：",idUserAuth)
	}
	u.ServeJSON()
}

/**
获取用户的展示图片组
 */
func (u *UserController) GetUserPic() {
	uid, _ := u.GetInt("uid")
	num, userpics := models.GetUserPicById(uid)
	fmt.Println(num)
	u.Data["json"] = userpics
	u.ServeJSON()
}
