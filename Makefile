.PHONY: deps build run  \
	fmt lint test install_deps clean coverage format 

BIN="./bin"
SRC=$(shell find . -name "*.go")
GOVERALLS_INSTALL=go install github.com/mattn/goveralls@latest
GOVERALLS_CMD=goveralls
GO_PACKAGES=./pkg/... ./qtool-api/... ./qtool-cli/...
TEST_SCRIPT=richgo test -v ${GO_PACKAGES}

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

default: all

all: fmt test

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: install_deps
	$(info ******************** running tests ********************)
	${TEST_SCRIPT}
	# richgo test -v ${GO_PACKAGES}

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

coverage:
	$(info ******************** running coverage ********************)
	${GOVERALLS_INSTALL}
	if [ "${COVERALLS_TOKEN}" ]; then ${TEST_SCRIPT} -coverprofile=c.out -covermode=count; ${GOVERALLS_CMD} -coverprofile=c.out -repotoken ${COVERALLS_TOKEN}; fi

build:
	$(info ******************** building ********************)
	go build -v -o ${BIN}/qtool-api ./qtool-api/main.go
	go build -v -o ${BIN}/qtool-cli ./qtool-cli/main.go

run-docker:
	$(info ******************** running qtool-api on docker container ********************)
	docker build -t qtum/qtool-api .
	docker run -d -p 8080:8080 qtum/qtool-api

clean:
	rm -rf $(BIN)