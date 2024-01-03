# Makefile

REGISTRY ?= registry-1.ict-mec.net:18443/qwen-cli
TAG ?=

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GOOS ?= $(shell go env GOHOSTOS)
GOARCH ?= $(shell go env GOARCH)
COMMIT_REF ?= $(shell git rev-parse --verify HEAD)
BOILERPLATE_DIR = $(shell pwd)/hack
GOPROXY ?= https://goproxy.cn,direct
GOVERSION ?= 1.21.4

APP ?=
ifeq ($(APP),)
apps = $(shell ls cmd)
else
apps = $(APP)
endif

RACE ?=
ifeq ($(RACE),on)
	race = "-race"
endif

TAG ?=
ifeq ($(TAG),)
	TAG = $(COMMIT_REF)
endif

.PHONY: test-all
test-all:
	ginkgo -r -v

.PHONY: fmt
fmt:
	gofmt -w pkg cmd

.PHONY: go-build
go-build:
	CGO_ENABLED=0 go build -o cmd/qwen/qwen-cli -ldflags "-X 'main.app.version.version=${TAG}'" cmd/qwen/main.go