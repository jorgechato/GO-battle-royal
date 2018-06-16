MAIN_VERSION:=$(shell git describe --always --abbrev=0 --tags || echo "0.1")
COMMIT:=$()
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
AUTHOR:=$(shell git log -n 1 --pretty=format:"%an")
GOENV:=CGO_ENABLED=0 GOOS=linux
LDFLAGS:=-ldflags "-X main.version=${MAIN_VERSION} -X main.author=${AUTHOR} -X main.tag=${VERSION}"

.PHONY: deploy

default: test build

test: unit cov linter

deploy: clean build-deploy build-docker

install_dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

install_dev:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

install_unit:
	go get -u github.com/jstemmer/go-junit-report

build:
	go build ${LDFLAGS} -a -o server manage.go

run:
	go run ${LDFLAGS} manage.go

linter:
	gometalinter.v2 --exclude=manage.go --checkstyle > report.xml

cov:
	go test -coverprofile=coverage.out ./...

unit:
	go test -v ./... | go-junit-report > test.xml

clean:
	rm -rf server *.out *.xml

build-deploy:
	${GOENV} go build -a -installsuffix cgo ${LDFLAGS} -o server .

build-docker:
	docker build -t battle .

run-docker:
	docker run -p 8080:8080 --name battle battle

clean-deploy: clean
	docker rmi -f battle
	docker rm battle
