MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
LDFLAGS:=-ldflags "-X github.com/b3bas/b3bas-up/src/config.Version=${VERSION}"

default: run

depends:
	dep ensure -v

test:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

cover: test
	@go tool cover -html=coverage-all.out

run:
	rm -rf banner.text
	@ln -sf ./bin/banner.txt ./banner.txt
	@go run ${LDFLAGS} app.go

build: clean
	@go build ${LDFLAGS} -a -o ./bin/b3bas-up app.go

clean:
	rm -rf ./bin/b3bas-up coverage.out coverage-all.out
