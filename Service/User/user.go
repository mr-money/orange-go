package User

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shockerli/cvt"
	"go-study/Config"
	"go-study/Library/Cache"
	"go-study/Library/Handler"
	"go-study/Model"
	"go-study/Repository/User"
	"time"
)

//
// FindById
// @Description: 根据用户id查询用户信息
// @param uint64 id user_id 用户id
// @return Model.User
//
func FindById(userInfo *Model.User, id uint64) *Model.User {

	//redis key 数据库名:表名
	idKey := Cache.SetKey(
		cvt.String(Config.GetFieldByName(Config.Configs.Web, "DB", "DbName")),
		Model.TableName,
		cvt.String(id),
	)

	userJson, _ := Cache.Redis.Get(Cache.Cxt, idKey).Result()
	if len(userJson) > 0 {
		//json转指定struct
		userInterface := Handler.JsonToStruct(userJson, userInfo)

		return userInterface.(*Model.User)
	}

	//查询数据库
	Model.UserModel().Take(&userInfo, id)
	if userInfo.ID > 0 {
		//插入redis缓存
		Cache.Redis.Set(
			Cache.Cxt,
			idKey, Handler.ToJson(userInfo),
			1*time.Hour,
		)
	}

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
func Register(user map[string]string) (Model.User, string, error) {
	var userInfo Model.User

	userInfo.Name = user["name"]
	userInfo.Uuid = uuid.NewV4()

	//密码加密
	userInfo.Password = Handler.HashAndSalt(user["password"])

	//创建用户
	User.Create(userInfo)

	//todo 自动登录
	token, err := Handler.ApiLoginToken(userInfo)
	if err != nil {
		return Model.User{}, "", err
	}

	return userInfo, token, nil
}

//
// Login
// @Description: 登录
// @param user
// @return Model.User
// @return error
//
func Login(user map[string]string) (Model.User, string, error) {
	var userInfo Model.User

	userInfo.Name = user["name"]

	//查询用户
	User.FindUserByModel(&userInfo)

	if userInfo.ID == 0 {
		return userInfo, "", errors.New("用户名或密码错误")
	}

	//检查密码
	if !Handler.ComparePasswords(userInfo.Password, user["password"]) {
		return userInfo, "", errors.New("用户名或密码错误")
	}

	//生成jwt
	token, err := Handler.ApiLoginToken(userInfo)
	if err != nil {
		return Model.User{}, "", err
	}

	return userInfo, token, nil
}
