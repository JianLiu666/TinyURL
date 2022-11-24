GIT_NUM ?= ${shell git rev-parse --short=6 HEAD}
BUILD_TIME ?= ${shell date +'%Y-%m-%d_%T'}

.PHONY: help init demo shutdown-all shutdown-server shutdown-benchmark restart-infra restart-logger restart-server restart-benchmark lint local unit-test integration-test build-image

help:
	@echo "Usage: make [commands]\n"
	@echo "Comands:"
	@echo "  init               初始化建置環境 (docker volume, build image, etc.)"
	@echo "  demo               透過 docker-compose 啟動所有服務 (主要系統, 壓力測試工具, 各項監控工具)"
	@echo "  shutdown-all       關閉 docker-cpmpose 所有服務"
	@echo "  shutdown-server    關閉 docker-compose 上的 API Server"
	@echo "  restart-infra      重啟 docker-compose 上的資料庫 (MySQL, Redis)"
	@echo "  restart-logger     重啟 docker-compose 上的日誌系統"
	@echo "  restart-server     重啟 docker-compose 上的 API Server"
	@echo "  restart-benchmark  重啟 docker-compose 上的壓力測試工具"
	@echo "  lint               執行 Go Linter (golangci-lint)"
	@echo "  local              本地執行 API Server"
	@echo "  unit-test          本地執行單元測試腳本"
	@echo "  integration-test   本地執行整合測試腳本"
	@echo "  build-image        建置 API Server 映像檔"

init:
	rm -rf deployments/data
	mkdir -p deployments/data/mysql
	mkdir -p deployments/data/prometheus deployments/data/grafana
	mkdir -p deployments/data/mongodb deployments/data/elasticsearch deployments/data/graylog/data deployments/data/graylog/journal
	mkdir -p deployments/data/locust
	
	go mod download
	go mod tidy
	
	make build-image

demo:
	docker-compose -f deployments/04.locust.yaml down -v
	docker-compose -f deployments/03.monitoring.yaml down -v
	docker-compose -f deployments/02.server.yaml down -v
	docker-compose -f deployments/01.logger.yaml down -v
	docker-compose -f deployments/00.infra.yaml down -v

	docker-compose -f deployments/00.infra.yaml up -d
	docker-compose -f deployments/01.logger.yaml up -d
	docker-compose -f deployments/02.server.yaml up -d
	docker-compose -f deployments/03.monitoring.yaml up -d
	docker-compose -f deployments/04.locust.yaml up -d

	docker ps -a

shutdown-all:
	docker-compose -f deployments/04.locust.yaml down -v
	docker-compose -f deployments/03.monitoring.yaml down -v
	docker-compose -f deployments/02.server.yaml down -v
	docker-compose -f deployments/01.logger.yaml down -v
	docker-compose -f deployments/00.infra.yaml down -v

shutdown-server:
	docker-compose -f deployments/02.server.yaml down -v

shutdown-benchmark:
	docker-compose -f deployments/04.locust.yaml down -v

restart-infra:
	docker-compose -f deployments/00.infra.yaml down -v
	docker-compose -f deployments/00.infra.yaml up -d

restart-logger:
	docker-compose -f deployments/01.logger.yaml down -v
	docker-compose -f deployments/01.logger.yaml up -d

restart-server:
	docker-compose -f deployments/02.server.yaml down -v
	docker-compose -f deployments/02.server.yaml up -d

restart-benchmark:
	docker-compose -f deployments/04.locust.yaml down -v

	rm -rf deployments/data/locust
	mkdir -p deployments/data/locust
	
	cp -r test/benchmark/*.py deployments/data/locust/
	cp -r deployments/locust/ deployments/data/locust/
	
	docker-compose -f deployments/04.locust.yaml up -d

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
