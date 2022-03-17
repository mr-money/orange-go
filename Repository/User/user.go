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
