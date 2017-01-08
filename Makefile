# binary name to kill/restart
PROGRAM = viral-api

# targets not associated with files
.PHONY: dependencies default build test coverage clean kill restart serve

# check we have a couple of dependencies
dependencies:
	@command -v fswatch --version >/dev/null 2>&1 || { printf >&2 "fswatch is not installed, please run: brew install fswatch\n"; exit 1; }

# default targets to run when only running `make`
default: dependencies test

install:
	go get -t -v ./...

# clean up
clean:
	go clean

# run formatting tool and build
build: dependencies clean
	go fmt
	go build

# run unit tests
test: dependencies
	go test -v ./...

# attempt to kill running server
kill:
	-@killall -9 $(PROGRAM) 2>/dev/null || true

# attempt to build and start server
restart:
	@make kill
	@make build; (if [ "$$?" -eq 0 ]; then (./${PROGRAM} &); fi)

# watch .go files for changes then recompile & try to start server
# will also kill server after ctrl+c
serve: dependencies
	@make restart
	@fswatch -o ./**/*.go tmpl/* | xargs -n1 -I{} make restart || make kill
