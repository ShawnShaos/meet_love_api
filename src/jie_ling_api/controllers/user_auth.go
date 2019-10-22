package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserAuthController struct {
	beego.Controller
}

// AppCode 来源于阿里云服务器（需购买） https://market.console.aliyun.com/imageconsole/index.htm?#/bizlist?_k=qlypqg
//https://market.aliyun.com/products/57000002/cmapi025518.html?spm=5176.2020520132.101.2.563a72184D3Wdf
const (
	APP_CODE         = "APPCODE 74aaf337c51f4429a91a1e6378d6688b"
	ID_CARD_AUTH_URL = "https://idenauthen.market.alicloudapi.com/idenAuthentication" //身份证验证地址
)

/**
身份证认证
 */
func (u *UserAuthController) IdCardAuth() {
	name := u.GetString("name") //姓名
	idNo := u.GetString("idno") //身份证号
	data := "idNo=" + idNo + "&name=" + name
	beego.Info(name, idNo)
	request, err := http.NewRequest("POST", ID_CARD_AUTH_URL, strings.NewReader(data))

	if err != nil {
		beego.Error(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", APP_CODE)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		beego.Error(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		beego.Error(err)
	}
	fmt.Println(string(body))
	u.Data["json"] = string(body)
	u.ServeJSON()
}
