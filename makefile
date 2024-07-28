.PHONY: build
build:
	go generate ./...
	go build ./cmd/server

.PHONY: run
run: build
	@./server

.PHONY: clean
clean:
	@rm -f server
	@rm -rf ui/build
