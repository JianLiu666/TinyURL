# TinyURL

## Goal

Create a shorter aliases for original URLs.

## High Level System Design

![image](./doc/image/architecture.svg)

## Todo

- APIs
  - [ ] [POST] {domain}/api/v1/create
  - [ ] [GET] {domain}/api/v1/heartbeat
  - [ ] [GET] {domain}/api/v1/urls
  - [ ] [GET] {domain}/{tinyurl}

- Mechanisms
  - [ ] Account
  - [ ] Rate Limiter

- CI
  - [ ] Lint
  - [ ] Unit Testing
  - [ ] Integration Testing

- Monitoring
  - [ ] ELK
  - [ ] Grafana

## References

- [MurMurHash](https://en.wikipedia.org/wiki/MurmurHash)