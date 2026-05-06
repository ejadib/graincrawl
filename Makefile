.PHONY: test vet tidy check build run smoke release-snapshot clean

VERSION ?= 0.0.0-dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w \
	-X github.com/vincentkoc/graincrawl/internal/buildinfo.Version=$(VERSION) \
	-X github.com/vincentkoc/graincrawl/internal/buildinfo.Commit=$(COMMIT) \
	-X github.com/vincentkoc/graincrawl/internal/buildinfo.Date=$(DATE)

test:
	GOWORK=off go test -count=1 ./...

vet:
	GOWORK=off go vet ./...

tidy:
	GOWORK=off go mod tidy

check: tidy vet test
	git diff --exit-code -- go.mod go.sum

build:
	GOWORK=off go build -trimpath -ldflags "$(LDFLAGS)" -o bin/graincrawl ./cmd/graincrawl

run:
	GOWORK=off go run ./cmd/graincrawl --help

smoke: build
	tmp="$$(mktemp -d)"; \
	cfg="$$tmp/config.toml"; \
	db="$$tmp/graincrawl.db"; \
	GRAINCRAWL_DB_PATH="$$db" ./bin/graincrawl --config "$$cfg" init --json; \
	./bin/graincrawl --config "$$cfg" metadata --json; \
	./bin/graincrawl --config "$$cfg" status --json; \
	./bin/graincrawl --config "$$cfg" tui --json; \
	./bin/graincrawl --config "$$cfg" snapshot create --out "$$tmp/snapshot" --json

release-snapshot:
	GOWORK=off goreleaser release --snapshot --clean --skip=publish

clean:
	rm -rf bin dist
