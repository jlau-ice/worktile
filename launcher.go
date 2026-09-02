package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const (
	FRONTEND_PORT = "28881"
	BACKEND_PORT  = "28882"
)

func main() {
	// 创建 logs 目录
	logDir := "./logs"
	os.MkdirAll(logDir, 0755)

	// 1. 检查并关闭前端端口上的旧服务
	if isPortInUse(FRONTEND_PORT) {
		log.Printf("端口 %s 已被占用，正在尝试关闭旧服务...", FRONTEND_PORT)
		killProcessOnPort(FRONTEND_PORT)
		time.Sleep(2 * time.Second)
	}

	// 2. 检查并关闭后端端口上的旧服务
	if isPortInUse(BACKEND_PORT) {
		log.Printf("端口 %s 已被占用，正在尝试关闭旧服务...", BACKEND_PORT)
		killProcessOnPort(BACKEND_PORT)
		time.Sleep(2 * time.Second)
	}

	// 3. 启动后端
	startService("go", []string{"run", "./cmd/server/main.go"}, ".", "backend.log")

	// 等待后端服务启动完成
	log.Printf("等待后端服务启动...")
	waitForPort(BACKEND_PORT, 30*time.Second)

	// 4. 启动前端
	// 注意：Windows 下执行 npm 建议用 npm.cmd
	startService("cmd", []string{"/c", "npm run dev"}, "./frontend", "frontend.log")
}

// isPortInUse 检查端口是否被占用
func isPortInUse(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true
	}
	ln.Close()
	return false
}

// killProcessOnPort 关闭占用指定端口的进程
func killProcessOnPort(port string) {
	// 使用 netstat 查找占用端口的进程 PID
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("netstat -ano | findstr :%s", port))
	output, err := cmd.Output()
	if err != nil {
		log.Printf("查找端口进程失败: %v", err)
		return
	}

	if len(output) == 0 {
		return
	}

	// 解析输出获取 PID 并终止进程
	// netstat 输出格式: TCP    0.0.0.0:28881    0.0.0.0:0    LISTENING    1234
	cmd = exec.Command("cmd", "/c", fmt.Sprintf("for /f \"tokens=5\" %%a in ('netstat -ano ^| findstr :%s ^| findstr LISTENING') do taskkill /F /PID %%a", port))
	if err := cmd.Run(); err != nil {
		log.Printf("关闭旧服务失败: %v", err)
	} else {
		log.Printf("已成功关闭端口 %s 上的旧服务", port)
	}
}

// waitForPort 等待指定端口变为可用状态（服务已启动）
func waitForPort(port string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if isPortInUse(port) {
			log.Printf("端口 %s 已就绪", port)
			return true
		}
		time.Sleep(500 * time.Millisecond)
	}
	log.Printf("等待端口 %s 超时", port)
	return false
}

func startService(name string, args []string, dir string, logFile string) {
	logPath := filepath.Join("./logs", logFile)
	f, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = f
	cmd.Stderr = f

	// 核心：设置 SysProcAttr 隐藏窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}

	cmd.Start()
	log.Printf("服务已启动: %s (目录: %s, 日志: %s)", name, dir, logFile)
	// 注意：这里用 Start 而不是 Run，这样程序启动后会立即返回执行下一个，不会阻塞
}
