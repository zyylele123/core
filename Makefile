# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep build client install clean gen

.DEFAULT_GOAL := build
DOSPROXY_PATH := onchain/eth/contracts/DOSProxy.sol
CONTRACTS_GOPATH := onchain/dosproxy
GENERATED_FILES := $(shell find $(CONTRACTS_GOPATH) -name '*.go')
ETH_CONTRACTS := onchain/eth/contracts
BOOT_CREDENTIAL := testAccounts/bootCredential

build: dep client

dep:
	dep ensure -vendor-only

client: gen
	go build -o client

client-docker:
	docker build -t dosnode -f Dockerfile .

install: dep client
	go install

updateSubmodule:
	test -f $(DOSPROXY_PATH) || git submodule update --init --recursive
	git submodule update --remote

gen: updateSubmodule
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosproxy --out $(CONTRACTS_GOPATH)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out $(CONTRACTS_GOPATH)/DOSProxy.go

clean:
	rm -f client
	rm -f $(GENERATED_FILES)
