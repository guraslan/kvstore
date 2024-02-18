.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
		go fmt .
vet: fmt
		go vet .
build: vet
		mkdir -p build
		go build -o build/kvstore cmd/main.go
run: build
		build/kvstore
clean:
		rm -rf build