//go:build windows
// +build windows

package curl

import (
	"fmt"
	"os"
)

func getVersion() (string, error) {

	version := os.Getenv("CURL_VERSION")
	if version == "" {
		return "", fmt.Errorf("CURL_VERSION env not exist")
	}

	return version, nil
}
