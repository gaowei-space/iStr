GO111MODULE=on

.PHONY: build
build:
	go build $(RACE) -o bin/iStr ./cmd/

.PHONY: test
test:
	go test $(RACE) ./...

.PHONY: package
package:
	bash ./scripts/package.sh

.PHONY: package-all
package-all:
	bash ./scripts/package.sh -p 'linux darwin windows'

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm bin/iStr