GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}

.PHONY: help init demo shutdown-all shutdown-server restart-infra restart-logger restart-server restart-benchmark lint local unit-test integration-test build-image

help:
	@echo "Usage: make [commands]\n"
	@echo "Comands:"
	@echo "  init               initial container volumes and download needed third-party modules."
	@echo "  demo               enable whole needed images in containers with docker-compose."
	@echo "  shutdown-all       shutdown all containers with docker-compose."
	@echo "  shutdown-server    "
	@echo "  restart-infra      "
	@echo "  restart-logger     "
	@echo "  restart-server     "
	@echo "  restart-benchmark  "
	@echo "  lint               run golang linter (golangci-lint)."
	@echo "  local              enable tinyurl server in local environment."
	@echo "  unit-test          run unit test in local environment."
	@echo "  integration-test   run integration test in local environment."
	@echo "  build-image        start to build tinyurl image."

init:
	rm -rf deployment/data
	mkdir -p deployment/data/mysql
	mkdir -p deployment/data/prometheus deployment/data/grafana
	mkdir -p deployment/data/mongodb deployment/data/elasticsearch deployment/data/graylog/data deployment/data/graylog/journal
	mkdir -p deployment/data/locust
	
	go mod download
	go mod tidy
	
	make build-image

demo:
	docker-compose -f deployment/04.locust.yaml down -v
	docker-compose -f deployment/03.monitoring.yaml down -v
	docker-compose -f deployment/02.server.yaml down -v
	docker-compose -f deployment/01.logger.yaml down -v
	docker-compose -f deployment/00.infra.yaml down -v

	docker-compose -f deployment/00.infra.yaml up -d
	docker-compose -f deployment/01.logger.yaml up -d
	docker-compose -f deployment/02.server.yaml up -d
	docker-compose -f deployment/03.monitoring.yaml up -d
	docker-compose -f deployment/04.locust.yaml up -d

	docker ps -a

shutdown-all:
	docker-compose -f deployment/04.locust.yaml down -v
	docker-compose -f deployment/03.monitoring.yaml down -v
	docker-compose -f deployment/02.server.yaml down -v
	docker-compose -f deployment/01.logger.yaml down -v
	docker-compose -f deployment/00.infra.yaml down -v

shutdown-server:
	docker-compose -f deployment/02.server.yaml down -v

restart-infra:
	docker-compose -f deployment/00.infra.yaml down -v
	docker-compose -f deployment/00.infra.yaml up -d

restart-logger:
	docker-compose -f deployment/01.logger.yaml down -v
	docker-compose -f deployment/01.logger.yaml up -d

restart-server:
	docker-compose -f deployment/02.server.yaml down -v
	docker-compose -f deployment/02.server.yaml up -d

restart-benchmark:
	docker-compose -f deployment/04.locust.yaml down -v

	rm -rf deployment/data/locust
	mkdir -p deployment/data/locust
	
	cp -r test/benchmark/*.py deployment/data/locust/
	cp -r deployment/locust/ deployment/data/locust/
	
	docker-compose -f deployment/04.locust.yaml up -d

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

build-image:
	docker build -t tinyurl:latest .
