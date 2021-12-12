APPNAME := $(shell basename $(CURDIR))

ifeq (,$(shell git describe --abbrev=0 --tags 2>/dev/null))
VERSION := v0.0.0
else
VERSION := $(shell git describe --abbrev=0 --tags)
endif

ifeq (,$(shell git rev-parse --short HEAD 2>/dev/null))
REVISION := unknown
else
REVISION := $(shell git rev-parse --short HEAD)
endif

LDFLAGS_APPNAME  := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(VERSION)"
LDFLAGS_REVISION := -X "main.Revision=$(REVISION)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static'

SRCS := $(shell find $(CURDIR) -type f -name '*.go')

.PHONY: all
all: build

.PHONY: build
build: $(APPNAME)
$(APPNAME): $(CURDIR)/bin/$(APPNAME)
$(CURDIR)/bin/$(APPNAME): $(SRCS)
	buf generate
	CGO_ENABLED=0 go generate $(LDFLAGS) ./...
	CGO_ENABLED=0 go build $(LDFLAGS) -o $@

.PHONY: install
install: $(SRCS)
	CGO_ENABLED=0 go install $(LDFLAGS)

.PHONY: run
run: $(CURDIR)/bin/$(APPNAME)
	$(CURDIR)/bin/$(APPNAME)

.PHONY: archive
archive: $(CURDIR)/bin/$(APPNAME).zip
$(CURDIR)/bin/$(APPNAME).zip: $(CURDIR)/bin/$(APPNAME)
	cd $(CURDIR)/bin && zip $@ $(APPNAME)

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: tools
tools:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install golang.org/x/tools/cmd/stringer
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: lint
lint:
	staticcheck

.PHONY: release
release:
ifneq ($(GITHUB_TOKEN),)
	goreleaser release --rm-dist
endif

.PHONY: snapshot
snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: clean
clean:
	rm -rf $(CURDIR)/bin
	rm -rf $(CURDIR)/dist
