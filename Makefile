
GO  := $(shell which go)

.PHONY: gomod-tidy go.mod standup

## go.mod: init go.mod if file does not exist
go.mod:
	$(GOMOD_ENV) $(GO) mod init $(PACKAGE_NAME)

## gomod-tidy: update dependencies when references added or removed
gomod-tidy:
	$(GO) mod tidy
	$(GO) mod download
	$(GO) mod vendor

## standup: brings up the api server
standup:
	cd server && go run main.go