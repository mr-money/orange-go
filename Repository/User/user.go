package User

import (
	"go-study/Model"
)

//
// FindById
// @Description: 根据id获取用户
// @param id
// @return userInfo
//
func FindById(id uint64) (userInfo Model.User) {

	Model.UserModel().Take(&userInfo, id)

	return
}

//
// Create
// @Description: 新增用户方法
// @param user
// @return uint64
//
func Create(user Model.User) uint64 {
	Model.UserModel().Create(&user)

	return user.ID
}
