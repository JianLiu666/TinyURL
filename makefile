GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}
CONFIG_PATH ?= $(CURDIR)/conf.d
CONFIG_FILE ?= env.yaml

build_infra:
	make build_docker
	docker-compose -f infra/docker-compose.yaml down -v
	docker-compose -f infra/docker-compose.yaml up -d
	docker ps -a

build_docker:
	docker build -t tinyurl:latest . 

local_run:
	go run main.go -f conf.d/env.yaml server

unit_test:
	go test -v ./...

integration_test:
	go run main.go integration