package curl

import (
	"time"
)

var (
	safe     bool
	waitTime = 5 * time.Second
	version  = "7.47.0"
)

func init() {
	current, err := getVersion()
	if err != nil {
		panic(err)
	}
	if err = supportVersion(current, version); err != nil {
		panic(err)
	}
}

// Init check options and curl command maximum wait time
func Init(isSafe bool, wait int) {
	safe = isSafe
	waitTime = time.Duration(wait) * time.Second
}
