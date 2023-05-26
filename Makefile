.PHONY: vicion vicion-cross evm all test clean
.PHONY: vicion-linux vicion-linux-386 vicion-linux-amd64 vicion-linux-mips64 vicion-linux-mips64le
.PHONY: vicion-darwin vicion-darwin-386 vicion-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.20.1
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

vicion:
	go run build/ci.go install ./cmd/vicion
	@echo "Done building."
	@echo "Run \"$(GOBIN)/vicion\" to launch vicion."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

vicion-cross: vicion-windows-amd64 vicion-darwin-amd64 vicion-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/vicion-*

vicion-linux: vicion-linux-386 vicion-linux-amd64 vicion-linux-mips64 vicion-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-*

vicion-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/vicion
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep 386

vicion-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/vicion
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep amd64

vicion-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/vicion
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep mips

vicion-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/vicion
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep mipsle

vicion-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/vicion
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep mips64

vicion-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/vicion
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/vicion-linux-* | grep mips64le

vicion-darwin: vicion-darwin-386 vicion-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/vicion-darwin-*

vicion-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/vicion
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-darwin-* | grep 386

vicion-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/vicion
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-darwin-* | grep amd64

vicion-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) -buildmode=mode -x --targets=windows/amd64 -v ./cmd/vicion
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/vicion-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
