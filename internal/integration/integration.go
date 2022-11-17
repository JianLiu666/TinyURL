package integration

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
	"tinyurl/internal/config"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"

	"github.com/fatih/color"
)

type tester struct {
	kvStore      kvstore.KvStore
	rdb          rdb.RDB
	serverConfig config.ServerOpts
}

func NewIntegrationTester(kvstore kvstore.KvStore, rdb rdb.RDB, serverConfig config.ServerOpts) *tester {
	return &tester{
		kvStore:      kvstore,
		rdb:          rdb,
		serverConfig: serverConfig,
	}
}

func (t *tester) Start() {
	casef(t.testcase1)
	casef(t.testcase2)
	casef(t.testcase3)
	casef(t.testcase4)
}

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
