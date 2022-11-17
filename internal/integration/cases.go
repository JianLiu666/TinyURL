package integration

import (
	"context"
	"fmt"
	"tinyurl/internal/config"
)

// 模擬使用者申請短網址且使用短網址跳轉
func (t *tester) testcase1() {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(context.TODO())

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_200)
	funcf(s, redirect_302)
}

// 模擬使用者使用不存在的短網址跳轉
func (t *tester) testcase2() {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(context.TODO())

	s := &session{
		origin: "",
		atlas:  "",
		tiny: fmt.Sprintf("%s%s/api/v1/jianliu",
			config.Env().Server.Domain,
			config.Env().Server.Port),
	}

	funcf(s, redirect_400)
}

// 模擬使用者對相同網址重複製作短網址
func (t *tester) testcase3() {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(context.TODO())

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_200)
	funcf(s, create_400)
}

// 模擬使用者對相同網址重複製作短網址
// Redis cache 過期時的情境
func (t *tester) testcase4() {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(context.TODO())

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_200)

	t.kvStore.FlushAll(context.TODO())

	funcf(s, create_400)
}
