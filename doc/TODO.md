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

- [ ] Account
- [ ] Rate Limiter
- [ ] Horizontal Scaling

## Testings

- [ ] Unit Testing
- [x] Integration Testing
- [x] Benchmark Testing
  - [x] using `Locust(O)`, `go-stress-testing(X)`
  - [x] dockerize

## Continuous Integration
- [ ] Github Actions
  - [x] Lint
  - [ ] Test

## Monitoring

- [x] Prometheus
- [x] Grafana
- [ ] Logging
  - [ ] Study `ELK`, `Graylog`
- [ ] Distributed Tracing
  - [ ] Study `OpenTracing`, `OpenTelemetry`, `Jaeger`, `Zipkin`
