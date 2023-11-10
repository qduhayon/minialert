.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: install_go
install_go: ## clean previous install of go, then download and install go 1.21.4 for linux
	sudo rm -rf /usr/local/go
	wget -4 https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
	rm go1.21.4.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin
	go version

.PHONY: clean
clean: ## clean previous builds
	@echo  "Clean previous builds..."
	@rm -rf bin

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build: tidy ## build minialert
	@echo "Building minialert..."
	go build -o bin/minialert

.PHONY: build_windows
build_windows: tidy ## cross-compile minialert for windows amd64
	@echo "Building minialert for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/minialert_win_amd64.exe

.PHONY: build_linux
build_linux: tidy ## cross-compile minialert for linux amd64
	@echo "Building minialert for Linux..."
	GOOS=linux GOARCH=amd64 go build -o bin/minialert_linux_amd64

.PHONY: all
all: clean build build_windows build_linux ## clean and build all targets