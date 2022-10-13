package User

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

func UserLog(name string) (string, error) {
	return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
}
