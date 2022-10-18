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

- [MurMurHash](https://en.wikipedia.org/wiki/MurmurHash)

### Monitoring

- [docker-compose 搭建 Prometheus+Grafana监控系统](https://www.cnblogs.com/qdhxhz/p/16325893.html)
- [Prometheus+Grafana+Go服务自建监控系统入门](https://www.xhyonline.com/?p=1492)

## Todo

- APIs
  - [x] `[POST] {domain}/api/v1/create`
  - [ ] `[GET] {domain}/api/v1/heartbeat`
  - [ ] `[GET] {domain}/api/v1/urls`
  - [x] `[GET] {domain}/api/v1/{tinyurl}`

- Mechanisms
  - [ ] Account
  - [ ] Rate Limiter

- Continuous Integration
  - [ ] Lint
  - [x] Unit Testing
  - [x] Integration Testing

- Monitoring
  - [x] Prometheus
  - [x] Grafana
  - [ ] ELK
  - [ ] Graylog
  - [ ] Distributed Tracing