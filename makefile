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
	@BUILDKIT_PROGRESS=plain docker build -t thingserver:latest -f docker/server.dockerfile .

docker-run:
	@docker run --rm --name thingserver -p3000:3000 thingserver:latest server -a 0.0.0.0 -p 3000

lint:
	@golangci-lint run

/usr/local/bin/upx:
	@echo "installing upx..."
	@bash scripts/install_upx.sh

install-upx: /usr/local/bin/upx

compress: install-upx bin/cli
	@upx --force-overwrite -5 -q bin/cli

.PHONY: install-ui build server dev docker-build docker-run lint install-upx compress
