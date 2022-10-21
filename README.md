# TinyURL

- [TinyURL](#tinyurl)
  - [Goal](#goal)
  - [How to use](#how-to-use)
  - [High Level System Design](#high-level-system-design)
  - [File Architecture](#file-architecture)
  - [References](#references)
    - [TinyURL](#tinyurl-1)
    - [Testing](#testing)
    - [Monitoring](#monitoring)
    - [Github Actions](#github-actions)

---

## Goal

- Create a shorter aliases for original URLs.
- Side project practice (implementations, write documentations, deployment and monitoring)

## How to use

Initial container volumes and download needed third-party modules for go.

```
make init
```

Start application by docker-compose.

```
make demo
```

## High Level System Design

![image](./doc/image/architecture.svg)

## File Architecture

```
TinyURL
 ├─ .github/         # includes github actions
 ├─ benchmark/       # includes benchmark testcases
 ├─ cmd/             # golang cli (cobra)
 ├─ conf.d/          # includes configuration files
 ├─ config/          # golang config manager (viper)
 ├─ doc/             # includes documentations (sequence, db schema, etc.)
 ├─ infra/           # includes docker-compose, mysql seed, etc.
 ├─ integration/     # includes integration testcases
 ├─ pkg/             # includes core modules (apis, storages, etc.)
 ├─ util/            # includes common modules (base converter, etc.)
 ├─ .gitattributes
 ├─ .gitignore
 ├─ .golangci.yaml   # golang linter settings
 ├─ dockerfile
 ├─ go.mod           # go mod files
 ├─ go.sum           # go mou files
 ├─ LICENSE
 ├─ main.go          # program entry point
 ├─ makefile         # cli tool
 └─ README.md
```

## References

### TinyURL

- [[Wiki] MurMurHash](https://en.wikipedia.org/wiki/MurmurHash)

### Testing

- [[Doc] Locust](https://docs.locust.io/en/stable/)

### Monitoring

- [[Blog] docker-compose 搭建 Prometheus+Grafana监控系统](https://www.cnblogs.com/qdhxhz/p/16325893.html)
- [[Blog] Prometheus+Grafana+Go服务自建监控系统入门](https://www.xhyonline.com/?p=1492)

### Github Actions

- [[Github] Marketplace/Actions/Run golangci-lint](https://github.com/marketplace/actions/run-golangci-lint)
- [[Github] github-actions-golang](https://github.com/mvdan/github-actions-golang)
- [[Github] Quickstart for GitHub Actions](https://docs.github.com/en/actions/quickstart)
