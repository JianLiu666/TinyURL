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
- [ ] Server Cache
  - [ ] `api/v1/create`
  - [ ] `api/v1/{redirect}`
- [ ] Rate Limiter
  - [ ] `api/v1/create` : by IP
  - [ ] `api/v1/{redirect}` : muti rules (per IP and global maximum)
- [ ] Precompute TinyURL
- [ ] Add APIs for registered user
- [ ] Horizontal Scaling

## Testings

- [ ] Unit Testing
- [x] Integration Testing
- [x] Benchmark Testing
  - [x] used `Locust(O)`, `go-stress-testing(X)`
  - [x] dockerize

## Continuous Integration
- [ ] Github Actions
  - [x] Lint
  - [ ] Test

## Monitoring

- [x] Prometheus
- [x] Grafana
- [ ] Logging
  - [x] used `logrus(O)`, `zap(X)`
  - [ ] Study `ELK`, `Graylog`
- [ ] Distributed Tracing
  - [ ] Study `OpenTracing`, `OpenTelemetry`, `Jaeger`, `Zipkin`
