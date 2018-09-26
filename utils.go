package blue

import (
	"fmt"
	"os"
	"time"
)

func DebugLog(e ...interface{}) {
	var logdata string
	for _, v := range e {
		logdata = logdata + " " + v.(string)
	}
	fmt.Fprintln(os.Stdout, time.Now().Format("2006/01/02 - 15:04:05"), logdata)
}
