package user

import (
	"domisep/db"
)

//CreateUser 插入 user 数据至数据库
func CreateUser(user User) {
	db.DB.Save(&user)
}

//GetUserDetailByName 根据用户名获取用户详情
// func GetUserDetailByName(name string) (user User) {
// 	user = db.DB.First(&User{Name: name})
// 	return user
// }

//GetUserDetailByID 根据用户 ID 获取用户详情
// func GetUserDetailByID(id string) (user User) {
// 	user = db.DB.First(&User{ID: id})
// 	return user
// }
