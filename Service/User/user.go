package User

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-study/Library/Handler"
	"go-study/Model"
	"go-study/Queue"
	QueueUser "go-study/Queue/Worker/Api/User"
	"go-study/Repository/User"
	"log"
)

//
// FindById
// @Description: 根据用户id查询用户信息
// @param uint64 id user_id 用户id
// @return Model.User
//
func FindById(userInfo *Model.User, id uint64) *Model.User {

	userInfo = User.FindById(userInfo, id)

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
	for i := 0; i < 5; i++ {
		go User.Create(userInfo)
	}

	//自动登录
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

// QueueTest 队列测试
func QueueTest(name string) string {
	// 注册任务
	err := Queue.Server.RegisterTask("userLog", QueueUser.UserLog)
	if err != nil {
		log.Println("reg task failed", err)
		return name
	}

	//task signature
	signature := &tasks.Signature{
		Name: "userLog",
		Args: []tasks.Arg{
			{
				Name:  "name",
				Type:  "string",
				Value: name,
			},
		},
	}

	//发送任务
	asyncResult, err := Queue.Server.SendTask(signature)
	if err != nil {
		log.Fatal(err)
	}

	//获取结果
	res, err := asyncResult.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("queue get res is %v\n", tasks.HumanReadableResults(res))
	return name
}
