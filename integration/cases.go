package integration

import (
	"fmt"
	"tinyurl/config"
)

type session struct {
	origin string
	atlas  string
	tiny   string
}

func Start() {
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
		tiny: fmt.Sprintf("%s%s/jianliu",
			config.Env().Server.Domain,
			config.Env().Server.Port),
	}

	funcf(s, redirect_BadRequest)
}
