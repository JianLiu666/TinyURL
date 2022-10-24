GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}

.PHONY: help init demo shutdown-all shutdown-server restart-infra restart-server lint local unit-test integration-test benchmark-up benchmark-down build-image

help:
	@echo "Usage: make [commands]\n"
	@echo "Comands:"
	@echo "  init               initial container volumes and download needed third-party modules."
	@echo "  demo               enable whole needed images in containers with docker-compose."
	@echo "  shutdown-all       shutdown all containers with docker-compose."
	@echo "  shutdown-server    "
	@echo "  restart-infra      "
	@echo "  restart-server     "
	@echo "  lint               run golang linter (golangci-lint)."
	@echo "  local              enable tinyurl server in local environment."
	@echo "  unit-test          run unit test in local environment."
	@echo "  integration-test   run integration test in local environment."
	@echo "  benchmark-up       "
	@echo "  benchmark-down     "
	@echo "  build-image        start to build tinyurl image."

init:
	rm -rf deployment/data
	mkdir -p deployment/data/mysql deployment/data/prometheus deployment/data/grafana deployment/data/locust
	go mod download
	go mod tidy
	make build-image

demo:
	docker-compose -f deployment/server.yaml down -v
	docker-compose -f deployment/infra.yaml down -v
	docker-compose -f deployment/infra.yaml up -d
	docker-compose -f deployment/server.yaml up -d
	docker ps -a

shutdown-all:
	docker-compose -f deployment/server.yaml down -v
	docker-compose -f deployment/infra.yaml down -v

shutdown-server:
	docker-compose -f deployment/server.yaml down -v

restart-infra:
	docker-compose -f deployment/infra.yaml down -v
	docker-compose -f deployment/infra.yaml up -d

restart-server:
	docker-compose -f deployment/server.yaml down -v
	docker-compose -f deployment/server.yaml up -d

lint:
	golangci-lint run

local:
# TODO: 應該要先檢查 env.yaml 是否存在
	go run main.go -f conf.d/env.yaml server

unit-test:
# TODO: 應該要先確認 server 是否已啟動
	go test -v ./...

integration-test:
# TODO: 應該要先確認 server 是否已啟動
	go run main.go integration

benchmark-up:
# TODO: 應該要先確認 server 是否已啟動
# locust -f ./benchmark/locustfile.py
	cp benchmark/locustfile.py deployment/data/locust/locustfile.py
	docker-compose -f deployment/locust.yaml up -d

benchmark-down:
	docker-compose -f deployment/locust.yaml down -v

build-image:
	docker build -t tinyurl:latest .
