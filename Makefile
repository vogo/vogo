.PHONY: build

lint:
	golangci-lint run

license-check:
	# go install github.com/vogo/license-header-checker/cmd/license-header-checker@latest
	license-header-checker -v -a -r apache-license.txt . go

format:
		goimports -w -l .
		go fmt ./...
		gofumpt -w .

test:
		go test ./... -v

build: license-check format lint test