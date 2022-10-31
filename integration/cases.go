package integration

import (
	"context"
	"fmt"
	"tinyurl/config"
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
	mysql.GetInstance().Raw("TRUNCATE TABLE urls;")
	redis.Init()
	redis.GetInstance().FlushAll(context.TODO())

	casef(testcase1)
	casef(testcase2)
}

// 模擬使用者申請短網址且使用短網址跳轉
func testcase1() {

	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	funcf(s, create_OK)
	funcf(s, redirect_OK)
}

// 模擬使用者使用不存在的短網址跳轉
func testcase2() {
	s := &session{
		origin: "",
		atlas:  "",
		tiny: fmt.Sprintf("%s%s/api/v1/jianliu",
			config.Env().Server.Domain,
			config.Env().Server.Port),
	}

	funcf(s, redirect_BadRequest)
}
