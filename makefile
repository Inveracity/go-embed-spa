.PHONY: build
build:
	go generate ./...
	go build -ldflags="-w -s" ./cmd/server

.PHONY: run
run: build
	@./server

.PHONY: clean
clean:
	@rm -f server
	@rm -rf ui/build

.PHONY: docker-build
docker-build:
	@docker build -t svelte-server:latest -f docker/server.dockerfile .

.PHONY: docker-run
docker-run:
	@docker run --rm --name svelte-server -p3000:3000 svelte-server:latest
