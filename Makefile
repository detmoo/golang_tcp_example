# Makefile for Go runtime environs
GOPATH=$(shell go env GOPATH)
GOBIN=$(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN=$(GOPATH)/bin
endif

develop:
	git config --global url."https://bowdata:$(AZURE_DEVOPS_PAT)@dev.azure.com/bowdata".insteadOf "https://bowdata@dev.azure.com/bowdata"
	go mod edit -replace bowdata.test.go_tcp_echo/pkg=../pkg
	go mod edit -replace bowdata.test.go_tcp_echo/cmd=../cmd
	go clean -i -cache -modcache
	go get -u ./...

lint:
	go mod edit -replace bowdata.test.go_tcp_echo/pkg=../pkg
	go mod edit -replace bowdata.test.go_tcp_echo/cmd=../cmd
	go vet -json ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest && $(GOPATH)/bin/staticcheck ./...

test:
	go mod edit -replace bowdata.test.go_tcp_echo/pkg=../pkg
	go mod edit -replace bowdata.test.go_tcp_echo/cmd=../cmd
	go clean -testcache
	go test ./...

setup_version:
	go mod init bowdata.test.go_tcp_echo

build:
	go clean -i -cache -modcache
	go mod tidy
	go build -o $(GOBIN)/bowdata/test/go_tcp_echo

release:
	go mod edit -dropreplace=bowdata.test.go_tcp_echo/pkg
	go mod edit -dropreplace=bowdata.test.go_tcp_echo/cmd

.PHONY: setup_version develop lint test build release