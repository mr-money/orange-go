package Handler

import (
	"os/exec"
	"runtime"
)

//
// OpenUrl
// @Description: 打开默认浏览器访问url
// @param url 浏览器打开地址
// @return error
//
func OpenUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd.exe"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
