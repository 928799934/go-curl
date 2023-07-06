package curl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func getVersion() (string, error) {
	//curl --version|grep curl | awk '{print $2}'

	// 创建命令的标准输出管道
	var b bytes.Buffer

	cmd1 := exec.Command("curl", "--version")
	cmd2 := exec.Command("grep", "curl")
	cmd3 := exec.Command("awk", "{print $2}")

	pipe1Reader, pipe1Writer := io.Pipe()
	pipe2Reader, pipe2Writer := io.Pipe()

	// 将 cmd1 的标准输出连接到管道写入端
	cmd1.Stdout = pipe1Writer

	// 将 cmd2 的标准输入连接到管道读取端
	cmd2.Stdin = pipe1Reader

	// 将 cmd1 的标准输出连接到管道写入端
	//cmd2.Stdout = &b
	cmd2.Stdout = pipe2Writer

	// 将 cmd2 的标准输入连接到管道读取端
	cmd3.Stdin = pipe2Reader
	//
	//cmd3.Stdout = pipe3Writer
	cmd3.Stdout = &b

	// 启动 cmd3 进程
	if err := cmd3.Start(); err != nil {
		return "", fmt.Errorf("cmd3.Start() error(%v)", err)
	}
	// 启动 cmd2 进程
	if err := cmd2.Start(); err != nil {
		return "", fmt.Errorf("cmd2.Start() error(%v)", err)
	}
	// 启动 cmd1 进程并等待它完成
	if err := cmd1.Start(); err != nil {
		return "", fmt.Errorf("cmd1.Start() error(%v)", err)
	}

	if err := cmd1.Wait(); err != nil {
		return "", fmt.Errorf("cmd1.Wait() error(%v)", err)
	}
	_ = pipe1Writer.Close()

	// 关闭管道写入端，通知 cmd2 输入已完成
	if err := cmd2.Wait(); err != nil {
		return "", fmt.Errorf("cmd2.Wait() error(%v)", err)
	}
	_ = pipe2Writer.Close()
	_ = pipe1Reader.Close()

	if err := cmd3.Wait(); err != nil {
		return "", fmt.Errorf("cmd3.Wait() error(%v)", err)
	}
	_ = pipe2Reader.Close()

	line, _, err := bufio.NewReader(&b).ReadLine()
	if err != nil {
		return "", fmt.Errorf("bufio.NewReader(&b).ReadLine() error(%v)", err)
	}
	return string(line), nil
}
