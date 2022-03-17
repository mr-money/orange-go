package User

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shockerli/cvt"
	"go-study/Model"
	"go-study/Repository/User"
)

//
// CreateUser
// @Description: 创建用户业务逻辑层
// @param user
// @return uint64
//
func CreateUser(user map[string]interface{}) uint64 {
	var userInfo Model.User

	userInfo.Name = cvt.String(user["name"])
	userInfo.Uuid = uuid.NewV4()

	return User.Create(userInfo)
}
