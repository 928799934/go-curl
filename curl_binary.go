package curl

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func ExecBinary(options []string, header http.Header, body []byte, url string) ([]byte, error) {
	if safe {
		if err := supportOptions(options); err != nil {
			return nil, err
		}
	}

	args := options
	for k, v := range header {
		for _, vv := range v {
			args = append(args, "-H", fmt.Sprintf("%s: %s", k, vv))
		}
	}

	if body != nil {
		md5Data := md5.Sum(body)
		name := path.Join(os.TempDir(), hex.EncodeToString(md5Data[:]))
		os.WriteFile(name, body, 0666)
		defer os.Remove(name)
		args = append(args, "--data-binary", "@"+name)
	}

	args = append(args, url)

	cmd := exec.Command("curl", args...)
	// 创建命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("cmd.StdoutPipe() error(%v)", err)
	}
	// 启动命令
	if err = cmd.Start(); err != nil {
		return nil, fmt.Errorf("cmd.Start() error(%v)", err)
	}

	var (
		buf bytes.Buffer
		ok  bool
	)
	ch := make(chan error, 1)
	ctx, _ := context.WithTimeout(context.Background(), waitTime)

	// 使用带缓冲的读取器读取输出
	go func(buf *bytes.Buffer, reader *bufio.Reader) {
		if _, err = io.Copy(buf, reader); err != nil {
			ch <- fmt.Errorf("io.Copy(buf, reader) error(%v)", err)
			return
		}
		close(ch)
	}(&buf, bufio.NewReader(stdout))

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err, ok = <-ch:
		if !ok {
			break
		}
		// 关闭管道
		close(ch)
		return nil, err
	}

	// 等待命令执行完毕
	if err = cmd.Wait(); err != nil {
		return nil, fmt.Errorf("cmd.Wait() error(%v)", err)
	}
	return buf.Bytes(), nil
}

func StreamBinary(options []string, header http.Header, body []byte, url string, fn func([]byte) error) error {
	if safe {
		if err := supportOptions(options); err != nil {
			return err
		}
	}
	args := options
	for k, v := range header {
		for _, vv := range v {
			args = append(args, "-H", fmt.Sprintf("%s: %s", k, vv))
		}
	}

	if body != nil {
		md5Data := md5.Sum(body)
		name := path.Join(os.TempDir(), hex.EncodeToString(md5Data[:]))
		os.WriteFile(name, body, 0666)
		defer os.Remove(name)
		args = append(args, "--data-binary", "@"+name)
	}

	args = append(args, url)

	cmd := exec.Command("curl", args...)
	// 创建命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("cmd.StdoutPipe() error(%v)", err)
	}
	// 启动命令
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("cmd.Start() error(%v)", err)
	}

	var ok bool
	ch := make(chan error, 1)
	ctx, _ := context.WithTimeout(context.Background(), waitTime)
	// 使用带缓冲的读取器读取输出
	go func(reader *bufio.Reader) {
		var (
			buf      bytes.Buffer
			line     []byte
			isPrefix bool
		)
		for {
			// 持续读取一行输出
			if line, isPrefix, err = reader.ReadLine(); err != nil {
				if err == io.EOF {
					break // 读取完毕
				}
				ch <- fmt.Errorf("reader.ReadLine() error(%v)", err)
				return // 出现错误时退出循环
			}
			buf.Write(line)
			if isPrefix {
				continue
			}
			if err = fn(buf.Bytes()); err != nil {
				break
			}
			buf.Reset()
		}
		close(ch)
	}(bufio.NewReader(stdout))

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err, ok = <-ch:
		if !ok {
			break
		}
		// 关闭管道
		close(ch)
		return err
	}

	// 等待命令执行完毕
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("cmd.Wait() error(%v)", err)
	}
	return nil
}
