package integration

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

// format testcase output
func casef(callback func()) {
	fmt.Printf("========== %s ==========\n", getFunctionName(callback))
	callback()
}

// format business logic output
func funcf(s *session, callback func(s *session) (bool, error)) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	start := time.Now()
	status, err := callback(s)
	elapsed := time.Since(start).Milliseconds()

	if status {
		fmt.Printf("| %3dms | %s\n", elapsed, green(getFunctionName(callback)))
	} else {
		fmt.Printf("| %3dms | %s : %s\n", elapsed, red(getFunctionName(callback)), err)
	}
}

func getFunctionName(f interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
}
