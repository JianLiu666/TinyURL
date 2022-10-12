GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}
CONFIG_PATH ?= $(CURDIR)/conf.d
CONFIG_FILE ?= env.yaml

local_run:
	go run main.go

unit_test:
	go test -v ./...