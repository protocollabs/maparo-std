TARGET := maparo
PACKAGE := github.com/protocollabs/maparo
DATE    := $(shell date +%FT%T%z)
VERSION := $(shell git describe --tags --always --dirty)
GOBIN   :=$(GOPATH)/bin

LDFLAGS  = "-X $(PACKAGE)/core.BuildVersion=$(VERSION) -X $(PACKAGE)/core.BuildDate=$(DATE)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


.PHONY: all build clean install uninstall fmt simplify check run

all: $(TARGET)

$(TARGET): $(SRC)
	go build -ldflags $(LDFLAGS) -o maparo cmd/maparo/maparo.go

install:
	go install -ldflags $(LDFLAGS) cmd/maparo/maparo.go

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

fmt:
	gofmt -l -w $(SRC)

simplify:
	gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}
