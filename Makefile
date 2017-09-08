PACKAGE = github.com/Axect/RGE
BASE = $(GOPATH)/src/$(PACKAGE)

default:
	glide update
	go build cmd/main.go
