package User

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shockerli/cvt"
	"go-study/Model"
	"go-study/Repository/User"
)

//
// FindUser
// @Description: 根据用户id查询用户信息
// @param uint64 id user_id 用户id
// @return Model.User
//
func FindUser(id uint64) Model.User {
	userInfo := Model.User{}
	User.FindById(&userInfo, id)

	return userInfo
}

//
// SelectUserListPage
// @Description: 分页获取用户列表
// @param search 搜索条件
// @param page 页数
// @param pageSize 每页条数
// @return []Model.User
//
func SelectUserListPage(search map[string]interface{}, page uint64, pageSize uint64) []Model.User {
	var userList []Model.User

	offset := (page - 1) * pageSize

	User.SelectPage(&userList, search, offset, pageSize)

	return userList

}

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
