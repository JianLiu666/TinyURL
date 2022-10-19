# TinyURL

## Goal

Create a shorter aliases for original URLs.

## How to use

Download third-party packages

```
go mod download
```

Start Service

```
make build_infra
```

## High Level System Design

![image](./doc/image/architecture.svg)

## References

### TinyURL

- [[Wiki] MurMurHash](https://en.wikipedia.org/wiki/MurmurHash)

### Monitoring

- [[Blog] docker-compose 搭建 Prometheus+Grafana监控系统](https://www.cnblogs.com/qdhxhz/p/16325893.html)
- [[Blog] Prometheus+Grafana+Go服务自建监控系统入门](https://www.xhyonline.com/?p=1492)

### Github Actions

- [[Github] Marketplace/Actions/Run golangci-lint](https://github.com/marketplace/actions/run-golangci-lint)
- [[Github] github-actions-golang](https://github.com/mvdan/github-actions-golang)
- [[Github] Quickstart for GitHub Actions](https://docs.github.com/en/actions/quickstart)

## Todo

- APIs
  - [x] `[POST] {domain}/api/v1/create`
  - [ ] `[GET] {domain}/api/v1/heartbeat`
  - [ ] `[GET] {domain}/api/v1/urls`
  - [x] `[GET] {domain}/api/v1/{tinyurl}`

- Mechanisms
  - [ ] Account
  - [ ] Rate Limiter
  - [ ] Horizontal Scaling

- Tests
  - [x] Unit Testing
  - [x] Integration Testing
  - [ ] Benchmark

- Continuous Integration
  - [ ] Github Actions
    - [x] Lint
    - [ ] Test

- Monitoring
  - [x] Prometheus
  - [x] Grafana
  - [ ] ELK
  - [ ] Graylog
  - [ ] Distributed Tracing