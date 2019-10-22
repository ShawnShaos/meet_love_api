package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetUserById(uid int) (UserBase, bool) {
	o := orm.NewOrm()
	u := UserBase{Id: uid}
	err := o.Read(&u)

	var userpicclub []* UserPic
	num, err := o.QueryTable("jie_lin_user_picclub").All(&userpicclub,"Img")

	if err == nil {
		fmt.Printf("数量", num)
		for _,v := range userpicclub{
			fmt.Println(v)
		}
	}else {
		fmt.Println(err)
	}
	return u,false
}
