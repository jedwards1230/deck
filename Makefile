.PHONY: run build test lint clean install

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE    ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS  = -X github.com/jedwards1230/deck/internal/version.Version=$(VERSION) \
           -X github.com/jedwards1230/deck/internal/version.Commit=$(COMMIT) \
           -X github.com/jedwards1230/deck/internal/version.Date=$(DATE)

run:
	go run . examples/slides.md

build:
	go build -ldflags "$(LDFLAGS)" -o deck .

install:
	go install -ldflags "$(LDFLAGS)" .

test:
	go test ./... -count=1

lint:
	golangci-lint run ./...

clean:
	rm -f deck
