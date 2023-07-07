# curl


# usage

```go

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
	data, _ := Http1Get(options, header, uri)
	fmt.Print(string(data))
}

```