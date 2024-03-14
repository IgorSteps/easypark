GO := go
WIRE := wire

CMD_DIR := ./cmd
DIST_DIR := ./dist
INTERNAL_DIR := ./internal
FUNCTIONAL_DIR := ./tests/functional

## clean-db: truncates db tables
.PHONY: clean-db
clean-db:
	./build/cleandb.sh

## build: build the project
.PHONY: build
build:
	$(GO) build -o $(DIST_DIR)/easypark $(CMD_DIR)/easypark

## run: run easypark 
.PHONY: run
run:
	./dist/easypark

## unit: runs unit tests and creates test coverage report.
.PHONY: unit
unit:
	$(GO) test $(INTERNAL_DIR)/...  -coverprofile=unit-test-coverage.out

## vendor: copy dependencies from Go to our repository.
.PHONY: vendor
vendor:
	$(GO) mod vendor

## tidy: tidy up mod file
.PHONY: tidy
tidy:
	$(GO) mod tidy

## mocks: generate mocks
.PHONY: mocks
mocks:
	rm -rf mocks
	mockery --dir=$(INTERNAL_DIR) --output=./mocks

## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/easypark

## coverage-report: generate test coverage report.
.PHONY: coverage-report
coverage-report:
	$(GO) tool cover -func=unit-test-coverage.out

## functional: runs functional tests
.PHONY: functional
functional:
	$(GO) test $(FUNCTIONAL_DIR)/...