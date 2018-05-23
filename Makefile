MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v vendor)
LDFLAGS:=-ldflags "-X github.com/jorgechato/battle-royale.Version=${VERSION}"

default: test build

test: unit cov linter

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
	gometalinter.v2 --exclude=version.go --checkstyle > report.xml

cov:
	set -e
	$(foreach pkg,$(PACKAGES), \
    	go test -coverprofile=coverage.out ${pkg};)

unit:
	go test -v ./... | go-junit-report > test.xml

clean:
	rm -rf server *.out *.xml
