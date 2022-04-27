package User

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-study/Library/Handler"
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
// Register
// @Description: 用户注册
// @param user
// @return uint64
//
func Register(user map[string]string) uint64 {
	var userInfo Model.User

	userInfo.Name = user["name"]
	userInfo.Uuid = uuid.NewV4()

	//密码加密
	userInfo.Password = Handler.HashAndSalt(user["password"])

	//todo 自动登录

	return User.Create(userInfo)
}

//
// Login
// @Description: 登录
// @param user
// @return Model.User
// @return error
//
func Login(user map[string]string) (Model.User, error) {
	var userInfo Model.User

	userInfo.Name = user["name"]

	//查询用户
	User.FindUserByModel(&userInfo)

	if userInfo.ID == 0 {
		return userInfo, errors.New("用户名或密码错误")
	}

	//检查密码
	if !Handler.ComparePasswords(userInfo.Password, user["password"]) {
		return userInfo, errors.New("用户名或密码错误")
	}

	// todo 生成jwt

	return userInfo, nil
}
