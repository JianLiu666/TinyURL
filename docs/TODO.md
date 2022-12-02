# TODO

- [TODO](#todo)
  - [APIs](#apis)
  - [Mechanisms](#mechanisms)
  - [Testings](#testings)
  - [Continuous Integration](#continuous-integration)
  - [Monitoring](#monitoring)

---

## APIs

- [x] `[POST] /api/v1/create`
- [ ] `[GET] /api/v1/heartbeat`
- [ ] `[GET] /api/v1/urls`
- [x] `[GET] /api/v1/{tinyurl}`

## Mechanisms

- [x] Set MySQL connection pool
- [x] Tuning MySQL configuration
- [x] Server Cache
  - [x] `api/v1/create`
  - [x] `api/v1/{redirect}`
- [x] Adapter Pattern (interface)
- [ ] Rate Limiter
  - [ ] `api/v1/create` : by IP
  - [ ] `api/v1/{redirect}` : muti rules (per IP and global maximum)
- [ ] Bloom Filter
- [ ] Precompute TinyURL
- [ ] Add APIs for registered user
- [ ] Horizontal Scaling
- [ ] Swagger

## Testings

- [ ] Unit Testing
- [x] Integration Testing
- [ ] Benchmark Testing
  - [x] used `Locust(O)`, `go-stress-testing(X)`
  - [x] dockerize
  - [ ] Mysql cases
    - [ ] 比較 primary key (number) 跟 unique key (string) 個別設立 index 時的寫入開銷差距
    - [ ] 比較 upsert 的執行開銷差距: replace(delete+insert) 與 insert on dunplicate(select+update)
  - [x] Encode tinyurl cases
    - [x] mermer3 與 sha256, md5 的開銷差距
  - [ ] 加入 Jaeger 前後的效能損失

## Continuous Integration
- [ ] Github Actions
  - [x] Lint
  - [ ] Test

## Monitoring

- [x] Prometheus
- [x] Grafana
- [x] Logging
  - [x] used `logrus(O)`, `zap(X)`
  - [x] used `Graylog(O)`, `ELK(X)`
- [x] Distributed Tracing (Application Performance Monitoring, APM)
  - [x] study `OpenTracing`, `OpenCensus`, `OpenTelemetry`
  - [x] used `Jaeger(O)`, `Zipkin(X)`, `SkyWalking(X)`
    - [x] Fiber
    - [x] Gorm
    - [x] Redis
- [ ] Pyroscope adhoc