#===================##===================##===================##===================##===================##===================##===================##===================#
#====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======#
#===================##===================##===================##===================##===================##===================##===================##===================#
BIN_MAIN = _main
BIN_TEST = _test
DONE_MESSAGE = "Finished Operation"

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

#===================##===================##===================##===================##===================##===================##===================##===================#
#====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

## up-dev: build up the realated containers
up-dev: up-main

## up-main: up docker for main service
up-main:
	@ echo "Up main service container"
	@ docker-compose -f project/src/services/main/docker/development/docker-compose.yml \
  --verbose up --build --remove-orphans -d
	@ echo ${DONE_MESSAGE}

#===================##===================##===================##===================##===================##===================##===================##===================#
#====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

.PHONY: test
# test: combine all tests through the entire project
test:
	@ echo "Running general tests..."
	@ rm -f ./out/${BIN_TEST} || true \
  && go test ./test/... -c -o ./out/${BIN_TEST} \
  && ./out/${BIN_TEST} -test.v -test.bench=.
	@ echo "Finish running general tests"

.PHONY: clean
## clean: clean binaries
clean:
	@ echo "Removing binaries"
	@ rm -r ./out || true
	@ echo ${DONE_MESSAGE}

.PHONY: build-main
## build-main: build de main service executable
build-main:
	@ rm -f ./out/${BIN_MAIN} || true \
	&& GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 \
	go build -o ./out/${BIN_MAIN} ./src/services/main/cmd/*.go

build-main-linux:
	@ rm -f ./out/${BIN_MAIN} || true \
	&& GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
	go build -o ./out/${BIN_MAIN} ./src/services/main/cmd/*.go

## run-main: run main service executable
run-main:
	@ ./out/${BIN_MAIN}

#===================##===================##===================##===================##===================##===================##===================##===================#
#====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

.PHONY: commit
## commit: add and commit files to develop branch
commit:
	@ cd .. && git add . && git commit

.PHONY: status
## status: show git status
status:
	@ cd .. && git status -u

.PHONY: push
## push: push application to main branch
push:
	@ cd .. && git push -u origin main

.PHONY: push-d
## push-d: push application to develop branch
push-d:
	@ cd .. && git push -u origin develop

