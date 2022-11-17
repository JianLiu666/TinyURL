package integration

import (
	"context"
	"fmt"
)

// 模擬使用者申請短網址且使用短網址跳轉
func (t *tester) testcase1(ctx context.Context) {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(ctx)

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, t.create_200)
	funcf(s, t.redirect_302)
}

// 模擬使用者使用不存在的短網址跳轉
func (t *tester) testcase2(ctx context.Context) {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(ctx)

	s := &session{
		origin: "",
		atlas:  "",
		tiny: fmt.Sprintf("%s%s/api/v1/jianliu",
			t.serverConfig.Domain,
			t.serverConfig.Port),
	}

	funcf(s, t.redirect_400)
}

// 模擬使用者對相同網址重複製作短網址
func (t *tester) testcase3(ctx context.Context) {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(ctx)

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, t.create_200)
	funcf(s, t.create_400)
}

// 模擬使用者對相同網址重複製作短網址
// Redis cache 過期時的情境
func (t *tester) testcase4(ctx context.Context) {
	t.rdb.Exec("TRUNCATE TABLE urls;")
	t.kvStore.FlushAll(ctx)

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, t.create_200)

	t.kvStore.FlushAll(ctx)

	funcf(s, t.create_400)
}
