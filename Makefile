ARTIFACT_NAME := external-dns-provider-adguard

TESTPARALLELISM := 4

WORKING_DIR := $(shell pwd)

.PHONY: lint
lint::
	golangci-lint run -c .golangci.yml
	go vet ./...

.PHONY: clean
clean::
	rm -r $(WORKING_DIR)/bin

.PHONY: build
build::
	go build -o $(WORKING_DIR)/bin/${ARTIFACT_NAME} ./cmd/webhook
	chmod +x $(WORKING_DIR)/bin/${ARTIFACT_NAME}

.PHONY: test
test::
	go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h -covermode atomic -coverprofile=covprofile ./...
