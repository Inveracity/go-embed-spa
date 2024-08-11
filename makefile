ui/node_modules:
	@cd ui && npm install
ui/dist:
	@go generate ./...

build:
	@go generate ./...
	@CGO_ENBALED=0 go build -ldflags="-w -s" -o ./bin/cli ./cmd/cli

server: ui/dist
	@go run cmd/cli/main.go server

dev: ui/node_modules
	@cd ui && npm run dev

docker-build:
	@docker build -t svelte-server:latest -f docker/server.dockerfile .

docker-run:
	@docker run --rm --name svelte-server -p3000:3000 svelte-server:latest server -p 3000 --apiport 3001

lint:
	@golangci-lint run

.PHONY: install-ui build server dev docker-build docker-run lint
