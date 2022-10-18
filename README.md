# TinyURL

## Goal

Create a shorter aliases for original URLs.

## High Level System Design

![image](./doc/image/architecture.svg)

## Todo

- APIs
  - [x] [POST] {domain}/api/v1/create
  - [ ] [GET] {domain}/api/v1/heartbeat
  - [ ] [GET] {domain}/api/v1/urls
  - [x] [GET] {domain}/api/v1/{tinyurl}

- Mechanisms
  - [ ] Account
  - [ ] Rate Limiter
  - [ ] Logger
  - [x] Monitoring

- CI
  - [ ] Lint
  - [x] Unit Testing
  - [x] Integration Testing

- Monitoring
  - [ ] ELK
  - [x] Grafana

## References

### TinyURL

- [MurMurHash](https://en.wikipedia.org/wiki/MurmurHash)

### Monitoring

- [docker-compose 搭建 Prometheus+Grafana监控系统](https://www.cnblogs.com/qdhxhz/p/16325893.html)
- [Prometheus+Grafana+Go服务自建监控系统入门](https://www.xhyonline.com/?p=1492)