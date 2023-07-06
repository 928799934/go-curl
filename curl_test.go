package curl

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_XX(t *testing.T) {
	uri := "https://www.baidu.com"

	options := []string{
		"--compressed", "--insecure",
	}
	header := http.Header{}
	header.Set("aa", "cc")
	data, _ := Http1Get(options, header, uri)
	fmt.Print(string(data))

	data, _ = Http1Post(options, header, []byte("hello"), uri)
	fmt.Print(string(data))
}
