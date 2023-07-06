package curl

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func supportOptions(options []string) error {
	options = append(options, "--data")

	cmd := exec.Command("curl", "--help", "all")
	// 创建命令的标准输出管道
	var b bytes.Buffer
	cmd.Stdout = &b
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cmd.Run() error(%v)", err)
	}
	help := b.String()
	for _, v := range options {
		if v[0] != '-' {
			continue
		}
		v = " " + v + " "
		if strings.Index(help, v) == -1 {
			return fmt.Errorf("curl option(%s) not support", v)
		}
	}
	return nil
}

func supportVersion(current, target string) error {
	var currentVersion []int
	for _, v := range strings.Split(current, ".") {
		i, _ := strconv.Atoi(v)
		currentVersion = append(currentVersion, i)
	}

	var targetVersion []int
	for _, v := range strings.Split(target, ".") {
		i, _ := strconv.Atoi(v)
		targetVersion = append(targetVersion, i)
	}

	currentVersionLen, targetVersionLen := len(currentVersion), len(targetVersion)
	switch {
	case currentVersionLen > targetVersionLen: // 7.61.1.2 vs 7.61.1 => 7.61.1.2 vs 7.61.1.0
		for i := currentVersionLen - targetVersionLen; i > 0; i-- {
			targetVersion = append(targetVersion, 0)
		}
	case currentVersionLen < targetVersionLen: // 7.61.1 vs 7.61.1.2 => 7.61.1.0 vs 7.61.1.2
		for i := targetVersionLen - currentVersionLen; i > 0; i-- {
			currentVersion = append(currentVersion, 0)
		}
	}

	for i, m := 0, len(targetVersion); i < m; i++ {
		if currentVersion[i] == targetVersion[i] {
			continue
		}

		if currentVersion[i] > targetVersion[i] {
			break
		}
		return fmt.Errorf("curl version must >= %s", version)
	}
	return nil
}
