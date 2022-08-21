.PHONY: deps build run run-api-docker \
	fmt lint test install_deps clean coverage format 
SHELL=/bin/bash
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

build-api:
	$(info ******************** building qtool-api ********************)
	go build -v -o ${BIN}/qtool-api ./qtool-api/main.go

build-cli:
	$(info ******************** building qtool-cli ********************)
	go build -v -o ${BIN}/qtool-cli ./qtool-cli/main.go

run-api:
	$(info ******************** running qtool-api on docker container ********************)
	@IMAGE=$$(docker images -q qtum/qtool-api 2> /dev/null) ; \
	if [ "$$IMAGE" == "" ] ; then docker build -t qtum/qtool-api . ; fi
	docker run -d --rm -p 8080:8080 --name qtool-api qtum/qtool-api

run-react:
	$(info ******************** running qtool react-web-app on docker container ********************)
	@IMAGE=$$(docker images -q qtum/qtool-react 2> /dev/null) ; \
	if [ "$$IMAGE" == "" ] ; then cd react-web-app; docker build -t qtum/qtool-react . ; fi
	docker run -d --rm -p 3000:80 --name qtool-react qtum/qtool-react

start-compose:
	$(info ******************** running qtool on docker compose ********************)
	@cd docker ; docker-compose up -d

stop-compose:
	$(info ******************** stopping containers from docker compose ********************)
	@cd docker ; docker-compose down

clean:
	rm -rf $(BIN)