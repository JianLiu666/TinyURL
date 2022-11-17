package integration

import (
	"context"
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
	ctx := context.Background()

	casef(ctx, t.testcase1)
	casef(ctx, t.testcase2)
	casef(ctx, t.testcase3)
	casef(ctx, t.testcase4)
}

// format testcase output
func casef(ctx context.Context, callback func(context.Context)) {
	fmt.Printf("========== %s ==========\n", getFunctionName(callback))
	callback(ctx)
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
