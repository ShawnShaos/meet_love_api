package filters

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"jie_ling_api/untils"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type UserJwtBase struct {
	Name string `json:name`
	Uid  int    `json:uid`
}

type UserJwt struct {
	UserJwtBase
	jwt.StandardClaims
}

var mySigningKey = []byte("AllYourBase")
/**
验证用户是否登录，基于token验证
 */
func Auth() beego.FilterFunc {
	return func(ctx *context.Context) {
		header := ctx.Request.Header
		tokenString := header.Get("Authorization")

		if len(tokenString) == 0 {
			err := ctx.Output.JSON(untils.GenTip(-1, "empty Authorization", ""), false, false)
			if err != nil {
				beego.Error("ctx.Output.JSON is err")
			}
			return
		} else {
			ret := CheckToken(tokenString)
			if ret["code"] == 1 { //检验token有效
				userJwtBase := ret["data"].(UserJwtBase) //interface转map
				ctx.Input.SetData("username", userJwtBase.Name) //用户名，昵称或真实姓名
				ctx.Input.SetData("uid", userJwtBase.Uid) //用户uid
				newtoken := RefreshToken(tokenString, userJwtBase)
				if newtoken["code"] == -1 {
					err := ctx.Output.JSON(newtoken, false, false)
					if err != nil {
						beego.Error("ctx.Output.JSON is err")
					}
					return
				}
				ctx.Output.Header("Authorization", newtoken["data"].(string))
			} else { //检验token无效
				err := ctx.Output.JSON(ret, false, false)
				if err != nil {
					beego.Error("ctx.Output.JSON is err")
				}
				return
			}
		}
	}
}

/**
	检查token是否有效
 */
func CheckToken(tokenString string) map[string]interface{} {

	token, err := jwt.ParseWithClaims(tokenString, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return untils.GenTip(-1, TokenMalformed.Error(), "")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return untils.GenTip(-1, TokenExpired.Error(), "")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return untils.GenTip(-1, TokenNotValidYet.Error(), "")
			} else {
				return untils.GenTip(-1, TokenInvalid.Error(), "")
			}
		}
	}
	if claims, ok := token.Claims.(*UserJwt); ok && token.Valid {
		return untils.GenTip(1, "ok",UserJwtBase{claims.Name,claims.Uid} )
}
return untils.GenTip(-1, "parse token fail", err.Error())
}

/**
获取token，采用map方式
 */
func CreatToken(userJwtBase UserJwtBase) map[string]interface{} {

	calim := UserJwt{
		userJwtBase,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 3000), //签发过期时间
			Issuer:    "admin",                         //签发者
			//IssuedAt:  int64(time.Now().Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)

	tokenSigned, err := token.SignedString(mySigningKey)

	if err != nil {
		return untils.GenTip(-1, "token signed fail", err)
	}
	return untils.GenTip(1, "ok", tokenSigned)
}

/*
刷新token
 */
func RefreshToken(tokenString string, userJwt UserJwtBase) map[string]interface{} {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	_, err := jwt.ParseWithClaims(tokenString, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return untils.GenTip(-1, err.Error(), "")
	}

	if ret := CreatToken(userJwt); ret["code"] == 1 { //刷新token成功
		return untils.GenTip(1, "RefreshToken is success", ret["data"])
	}

	return untils.GenTip(-1, TokenInvalid.Error(), "")
}
