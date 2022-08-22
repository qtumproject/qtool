.PHONY: deps build run run-api-docker \
	fmt lint test install_deps clean coverage format \
	stop-compose-dev
SHELL=/bin/bash
REBUILD=false
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

build-cli:
	$(info ******************** building qtool-cli ********************)
	go build -v -o ${BIN}/qtool-cli ./qtool-cli/main.go

build-docker-api:
	@IMAGE=$$(docker images -q qtum/qtool-api 2> /dev/null) ; \
	if [ "$$IMAGE" == "" || $(REBUILD) ] ; \
	$(info ******************** rebuilding qtool-api docker image ********************) \
	then docker build -t qtum/qtool-api . ; \ 
	fi

build-docker-react-prod:
	@IMAGE=$$(docker images -q qtum/qtool-react 2> /dev/null) ; \
	if [ "$$IMAGE" == "" || $(REBUILD) ] ; \
	$(info ******************** building qtool-react docker image ********************) \
	then cd react-web-app; docker build -t qtum/qtool-react:latest . ; \ 
	fi

# run-docker-api: build-docker-api
# 	$(info ******************** running qtool-api on docker container ********************)
# 	docker run -d --rm -p 8080:8080 --name qtool-api qtum/qtool-api

# run-docker-react: build-docker-react-prod
# 	docker run -d --rm -p 3000:80 --name qtool-react qtum/qtool-react

start-compose-prod:
	$(info ******************** running PROD qtool web on docker compose ********************)
	@cd docker ; docker-compose -f docker-compose.yml up -d

stop-compose-prod:
	$(info ******************** stopping PROD qtool web containers from docker compose ********************)
	@cd docker ; docker-compose -f docker-compose.yml down

start-compose-dev:
	$(info ******************** running DEV qtool web on docker compose ********************)
	@cd docker ; docker-compose -f docker-compose-dev.yml up -d

stop-compose-dev:
	$(info ******************** stopping DEV qtool web containers from docker compose ********************)
	@cd docker ; docker-compose -f docker-compose-dev.yml down

clean:
	rm -rf $(BIN)
	docker rm -f qtool-api qtool-react
	docker rmi -f qtum/qtool-api qtum/qtool-react