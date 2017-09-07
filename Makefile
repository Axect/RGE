PACKAGE = github.com/Axect/RGE
BASE = $(GOPATH)/src/$(PACKAGE)

default:
	go build cmd/main.go
