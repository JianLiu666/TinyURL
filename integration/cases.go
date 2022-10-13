package integration

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

type session struct {
	origin string
	atlas  string
	tiny   string
}

func Case1() {
	s := &session{
		origin: "https://tinyurl.com/app/",
		atlas:  "",
		tiny:   "",
	}

	middleware(s, create_ok)
	middleware(s, redirect_ok)
}

func middleware(s *session, callback func(s *session)) {
	start := time.Now()
	callback(s)
	log.Printf("| %4dms | %s",
		time.Since(start).Milliseconds(),
		getFunctionName(callback),
	)
}

func getFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
