package User

import (
	"go-study/Model"
)

func FindById(id uint64) (userInfo Model.User) {

	Model.UserModel.Where("id = ?", id).First(&userInfo)

	return
}
