package models

import (
	"fmt"
)

//`gorm:"type:bigint(20);column:id;primary_key"`
type User struct {
	UserId     int64  `json:"userId" gorm:"type:bigint(20);column:userId;primary_key"`
	Username   string `json:"username" gorm:"type:varchar(255);column:userName"`
	Gender     string `json:"gender" gorm:"type:varchar(255);column:gender"`
	Usermobile string `json:"usermobile" gorm:"type:varchar(255);column:userMobile"`
	Pwd        string `json:"pwd" gorm:"type:varchar(255);column:pwd"`
	Permission string `json:"permission" gorm:"type:varchar(255);column:permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(phone string, pwd string) error {
	if CheckUser(phone) {
		return fmt.Errorf("用户已存在！")
	}
	user := User{
		Usermobile: phone,
		Username:   phone,
		Pwd:        pwd,
		Gender:     "0",
	}
	if user.Usermobile == "18729580703" {
		user.Permission = "1"
	} else {
		user.Permission = "0"
	}
	db.Table("user").Create(&user)

	return nil
}

// CheckUser 检查用户是否存在
func CheckUser(phone string) bool {
	user := make([]*User, 0)
	db.Table("user").Where("userMobile=?", phone).Find(&user)

	if len(user) != 0 {
		return true
	}
	return false
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {
	user := make([]*User, 0)
	resultBool := false
	fmt.Println(loginReq)
	db.Table("user").Where("userMobile=?", loginReq.Phone).Where("pwd=?", loginReq.Pwd).Find(&user)
	if len(user) != 0 {
		resultBool = true
	}
	fmt.Println(user[0])
	return resultBool, *user[0], nil
}

// EditUserReq 更新用户信息数据类
type EditUserReq struct {
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	UserGender string `json:"gender"`
}

// UpdateUser 更新用户信息
func UpdateUser(editUser EditUserReq) (User, error) {
	user := new(User)
	user.Username = editUser.UserName
	user.Gender = editUser.UserGender
	db.Table("user").Where("userId=?", editUser.UserId).Update(&user)

	return *user, nil
}

//ResetPwd 重置密码
func ResetPwd(mobile string, pwd string) error {
	if !CheckUser(mobile) {
		return fmt.Errorf("用户不存在！")
	}
	user := new(User)
	user.Pwd = pwd
	db.Table("user").Where("userMobile=?", mobile).Update(&user)
	return nil
}
