package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type UserInfo struct {
	Id       int
	Name     string
	Password string
}

var (
	UserList map[string]*User
	o        orm.Ormer
)

func init() {
	o = orm.NewOrm()

	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

/*
添加用户
 */
func AddUserInfo(userinfo UserInfo) (int64, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := "SELECT id FROM user_info where name ='" + userinfo.Name + "'"
	num, _ := o.Raw(sql).Values(&maps)
	if num > 0 {
		return 0, errors.New("用户名已存在")
	}
	return o.Insert(&userinfo)
}

func GetUserPicById(uid int) (int64, []*UserPic) {
	o := orm.NewOrm()
	var pic []*UserPic
	num, err := o.QueryTable("jie_ling_user_pic").Filter("user_id", uid).All(&pic, "Img")
	if err == nil {
		return num, pic
	}
	return 0, nil
}

/**
根据opendId查询是否有授权用户
 */
func GetUserByOpenId(userAuth UserAuths) (map[string]interface{}, error) {
	err := o.Read(&userAuth, "Identifier")
	result := make(map[string]interface{})
	if err != nil { //无此用户
		return result,errors.New("用户不存在")
	} else { //存在，查询用户基本信息
		userBase := UserBase{Id: userAuth.UserBase.Id}
		o.Read(&userBase)
		result["uid"] = userBase.Id  //用户id
		result["nickname"] = userBase.Nickname  //用户昵称
 	return result, nil
}
}

/**
插入基础用户表和授权表
 */
func InsertUserBase(userbase UserBase) (int64) {
	id, err := o.Insert(&userbase) //基础用户表
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println("userbase插入错误：", err)
	}
	return id
}

/**
插入基础用户表和授权表
 */
func InsertUserAuths(userAuth UserAuths) (int64) {
	id, err := o.Insert(&userAuth) //基础用户表
	if err == nil {
		fmt.Println(id)
	}
	return id
}
