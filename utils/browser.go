package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	// 根据操作系统选择相应的命令
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", url)
	case "darwin": // MacOS
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported OS")
	}

	// 执行命令打开浏览器
	return cmd.Start()
}
