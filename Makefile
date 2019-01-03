MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
LDFLAGS:=-ldflags "-X github.com/b3bas/b3bas-up/src/config.Version=${VERSION}"

default: run

depends:
	@echo "Install and/or update dependencies..."
	@dep ensure -v
	@echo "Dependencies updated... [DONE]"

test:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

cover: test
	@echo "Running b3bas-up (test)..."
	@go tool cover -html=coverage-all.out

run:
	@echo "Running b3bas-up..."
	@rm -rf banner.text
	@ln -sf ./bin/banner.txt ./banner.txt
	@go run ${LDFLAGS} app.go

build: clean
	@go build ${LDFLAGS} -a -o ./bin/b3bas-up app.go
	@echo "Build binary...         [DONE]"

clean:
	@rm -rf ./bin/b3bas-up coverage.out coverage-all.out
	@echo "Cleanup binary...       [DONE]"
