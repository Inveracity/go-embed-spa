ui/dist:
	cd ./ui && npm run build

.PHONY: build
build: ui/dist
	go build ./cmd/server

.PHONY: clean
clean:
	@rm -f server
	@rm -rf ui/dist

.PHONY: run
run: build
	@./server
