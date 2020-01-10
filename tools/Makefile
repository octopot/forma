# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = install

SHELL = /bin/bash -euo pipefail

GO          = GOPRIVATE=$(GOPRIVATE) GOFLAGS=$(GOFLAGS) go
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   = go.octolab.net
GOPROXY     = direct
LOCAL       = $(MODULE)
MODULE      = $(shell $(GO) list -m)
PACKAGES    = $(shell $(GO) list ./... 2> /dev/null)
PATHS       = $(shell $(GO) list ./... 2> /dev/null | sed -e "s|$(MODULE)/\{0,1\}||g")
TIMEOUT     = 1s

ifeq (, $(PACKAGES))
	PACKAGES = $(MODULE)
endif

ifeq (, $(PATHS))
	PATHS = .
endif

.PHONY: go-env
go-env:
	@echo "GO111MODULE: $(shell $(GO) env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell $(GO) env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell $(GO) env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell $(GO) env GOPROXY))"
	@echo "LOCAL:       $(LOCAL)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"

.PHONY: deps
deps:
	@$(GO) mod tidy
	@if [[ "$(shell $(GO) env GOFLAGS)" =~ -mod=vendor ]]; then $(GO) mod vendor; fi

.PHONY: deps-clean
deps-clean:
	@$(GO) clean -modcache

.PHONY: update
update: selector = '.Require[] | select(.Indirect != true) | .Path'
update:
	@if command -v egg > /dev/null; then \
		packages="$(shell egg deps list)"; \
		$(GO) get -mod= -u $$packages; \
	elif command -v jq > /dev/null; then \
		packages="$(shell $(GO) mod edit -json | jq -r $(selector))"; \
		$(GO) get -mod= -u $$packages; \
	else \
		packages="$(shell cat go.mod | grep -v '// indirect' | grep '\t' | awk '{print $$1}')"; \
		$(GO) get -mod= -u $$packages; \
	fi

BINPATH = $(shell dirname $(PWD))/bin
LOCAL   = $(shell dirname $(MODULE))

export PATH := $(BINPATH):$(PATH)

.PHONY: tools-env
tools-env:
	@echo "BINPATH:     $(BINPATH)"

.PHONY: build
build: install

.PHONY: build-clean
build-clean:
	@$(GO) clean -cache

.PHONY: install
install:
	@ROOT=$(dir $(BINPATH)) $(GO) generate tools.go

.PHONY: install-clean
install-clean:
	@if command -v egg > /dev/null; then \
		tools="$(shell egg tools list)"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	else \
		tools="$(shell cat tools.go | grep 'go:generate' | awk '{print $$NF}' | xargs basename)"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	fi


.PHONY: clean
clean: build-clean deps-clean install-clean

.PHONY: env
env: go-env tools-env

.PHONY: refresh
refresh: update deps install
