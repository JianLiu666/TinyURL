GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}
CONFIG_PATH ?= $(CURDIR)/conf.d
CONFIG_FILE ?= env.yaml

build_infra:
	docker-compose -f infra/docker-compose.yaml down -v
	docker-compose -f infra/docker-compose.yaml up -d
	docker ps -a

local_run:
	go run main.go server

unit_test:
	go test -v ./...

integration_test:
	go run main.go integration