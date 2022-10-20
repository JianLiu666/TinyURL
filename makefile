GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}

.PHONY: help demo lint unit-test integration-test benchmark build-image

help:
	@echo "Usage: make [commands]\n"
	@echo "Comands:"
	@echo "  demo               enable whole needed images in containers with docker-compose."
	@echo "  lint               run golang linter (golangci-lint)."
	@echo "  server             enable tinyurl server in local environment."
	@echo "  unit-tst           run unit test in local environment."
	@echo "  integration-test   run integration test in local environment."
	@echo "  benchmark          run benchmark in local environment."
	@echo "  build-image        start to build tinyurl image."

demo:
	make build-image
	docker-compose -f infra/docker-compose.yaml down -v
	docker-compose -f infra/docker-compose.yaml up -d
	docker ps -a

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
	go run main.go benchmark
 
build-image:
	docker build -t tinyurl:latest .
