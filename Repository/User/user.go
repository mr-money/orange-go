package User

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shockerli/cvt"
	"go-study/Config"
	"go-study/Library/Cache"
	"go-study/Library/Handler"
	"go-study/Model"
	"reflect"
	"time"
)

//
// FindById
// @Description: 根据id获取用户
// @param id
// @param userInfo
//
func FindById(UserInfo *Model.User, Id uint64) *Model.User {

	//redis key 数据库名:表名
	idKey := Cache.SetKey(
		Config.GetFieldByName(Config.Configs.Web, "DB", "DbName"),
		Model.User{}.TableName(),
		cvt.String(Id),
	)

	userJson := Cache.RememberString(idKey, func() string {

		//避免闭包二次调用查询
		if UserInfo.ID == 0 {
			//查询数据库
			Model.UserModel().Take(&UserInfo, Id)
		}

		//默认为空
		userinfo := ""
		if !reflect.DeepEqual(UserInfo, Model.User{}) {
			userinfo = Handler.ToJson(UserInfo)
		}

		return userinfo

	}, 1*time.Hour)

	//json转指定struct
	Handler.JsonToStruct(userJson, &UserInfo)

	return UserInfo
}

//
// SelectPage
// @Description: 分页查询用户列表
// @param userList *[]Model.User 用户列表
// @param search 搜索条件
// @param offset 偏移量
// @param limit 每页条数
// @return *[]Model.User
//
func SelectPage(userList *[]Model.User, search map[string]interface{}, offset uint64, limit uint64) *[]Model.User {

	userModel := Model.UserModel().
		Where(&Model.User{Uuid: uuid.FromStringOrNil(cvt.String(search["uuid"]))}). //uuid搜索
		Where("name like ?", cvt.String(search["user_name"])+"%").                  //用户名模糊搜索
		Offset(cvt.Int(offset)).Limit(cvt.Int(limit))                               //分页参数

	userModel.Order("id DESC").Find(&userList)

	return userList
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

//
// FindUserByModel
// @Description: 根据结构体model查询用户
// @param user 用户model结构体
// @return *Model.User
//
func FindUserByModel(user *Model.User) *Model.User {
	Model.UserModel().Where(user).Take(&user)

	return user
}
