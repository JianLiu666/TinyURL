GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}

.PHONY: help init demo shutdown lint unit-test integration-test benchmark build-image

help:
	@echo "Usage: make [commands]\n"
	@echo "Comands:"
	@echo "  init               initial container volumes and download needed third-party modules."
	@echo "  demo               enable whole needed images in containers with docker-compose."
	@echo "  shutdowm           shutdown all containers with docker-compose."
	@echo "  lint               run golang linter (golangci-lint)."
	@echo "  server             enable tinyurl server in local environment."
	@echo "  unit-tst           run unit test in local environment."
	@echo "  integration-test   run integration test in local environment."
	@echo "  benchmark          run benchmark in local environment."
	@echo "  build-image        start to build tinyurl image."

init:
	rm -rf infra/data
	mkdir -p infra/data/mysql infra/data/prometheus infra/data/grafana
	go mod download
	go mod tidy
	make build-image

demo:
	docker-compose -f infra/docker-compose.yaml down -v
	docker-compose -f infra/docker-compose.yaml up -d
	docker ps -a

shutdown:
	docker-compose -f infra/docker-compose.yaml down -v

lint:
	golangci-lint run

server:
# TODO: 應該要先檢查 env.yaml 是否存在
	go run main.go -f conf.d/env.yaml server

unit-test:
# TODO: 應該要先確認 server 是否已啟動
	go test -v ./...

integration-test:
# TODO: 應該要先確認 server 是否已啟動
	go run main.go integration

benchmark:
# TODO: 應該要先確認 server 是否已啟動
	locust -f ./benchmark/locustfile.py
 
build-image:
	docker build -t tinyurl:latest .
