package Api

import (
	"github.com/shockerli/cvt"
	"orange-go/Service/User"
	"testing"
	"time"
)

// BenchmarkAddUser
// @Description: 性能测试 - 添加单个用户
// @param b
func BenchmarkAddUser(b *testing.B) {
	userInfo := make(map[string]string)
	userInfo["password"] = "123456"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userInfo["name"] = "bench-name" + cvt.String(time.Now().UnixNano())
		_, _, err := User.Register(userInfo)
		if err != nil {
			b.Errorf("Failed to register user: %v", err)
		}
	}
}

// TestAddUser999
// @Description: 测试批量添加用户（原AddUser999功能）
// @param t
func TestAddUser999(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestAddUser999 in short mode")
	}

	userInfo := make(map[string]string)
	userInfo["password"] = "123456"
	userInfo["name"] = "test-name" + cvt.String(time.Now().Unix())

	_, _, err := User.Register(userInfo)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}
}
