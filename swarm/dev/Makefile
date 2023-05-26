.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/vicion/vicion/cmd/swarm
	go build -o bin/vicion     github.com/vicion/vicion/cmd/vicion
	go build -o bin/bootnode github.com/vicion/vicion/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/vicion/vicion/swarm/...
