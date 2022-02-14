REGISTRY=reg.cvcio.org
REPO=github.com/cvcio/proto

BUF_VERSION:=1.0.0-rc6

generate:
	buf generate
build:
	buf build
update:
	buf mod update
lint:
	buf lint

install:
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"