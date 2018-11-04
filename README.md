# go-micro service template

This project aims to provide an enhanced template for go-micro services. Goal is to write code faster and not get distracted writing middleware code.

## Features
- [ ] Include _Example_ service protobuf
- [ ] Include _Makefile_
- [ ] Base handler including test
- [ ] Well structured main.go
- [ ] Configuration using **go-config** to provide configuration from various sources
- [ ] Pre-configured structured logging ready for parsing
- [ ] Log wrapper
- [ ] Prometheus metrics
- [ ] Trace wrapper
- [ ] Generator script

## Getting started
All you need to do is to checkout this repo and start hacking. But for the sake of simplicity let's just walk through all steps required to customize the template.

1. Adjust `DOCKER_IMAGE` and `DOCKER_TAG` in the **Makefile**
2. Rename the `example.proto` definition
3. Adjust the `proto` command in the *Makefile* to the new path