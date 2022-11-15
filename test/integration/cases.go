package integration

import (
	"context"
	"fmt"
	"tinyurl/internal/config"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/pkg/storage/redis"
)

type session struct {
	origin string
	atlas  string
	tiny   string
}

func Start() {
	mysql.Init()
	redis.Init()

	casef(testcase1)
	casef(testcase2)
	casef(testcase3)
	casef(testcase4)
}

// 模擬使用者申請短網址且使用短網址跳轉
func testcase1() {
	mysql.GetInstance().Exec("TRUNCATE TABLE urls;")
	redis.GetInstance().FlushAll(context.TODO())

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_200)
	funcf(s, redirect_302)
}

// 模擬使用者使用不存在的短網址跳轉
func testcase2() {
	mysql.GetInstance().Exec("TRUNCATE TABLE urls;")
	redis.GetInstance().FlushAll(context.TODO())

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
func testcase3() {
	mysql.GetInstance().Exec("TRUNCATE TABLE urls;")
	redis.GetInstance().FlushAll(context.TODO())

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
func testcase4() {
	mysql.GetInstance().Exec("TRUNCATE TABLE urls;")
	redis.GetInstance().FlushAll(context.TODO())

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_200)

	redis.GetInstance().FlushAll(context.TODO())

	funcf(s, create_400)
}
