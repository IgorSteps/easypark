GO := go
WIRE := wire

CMD_DIR := ./cmd
DIST_DIR := ./dist
INTERNAL_DIR := ./internal

## setup-db: truncates db tables
.PHONY: setup-db
setup-db:
	./build/cleandb.sh

## build: build the project
.PHONY: build
build:
	$(GO) build -o $(DIST_DIR)/easypark $(CMD_DIR)/easypark

## run: run easypark 
.PHONY: run
run:
	./dist/easypark

## unit: runs unit tests
.PHONY: unit
unit:
	$(GO) test $(INTERNAL_DIR)/...

## vendor: copy dependencies from Go to our repository.
.PHONY: vendor
vendor:
	$(GO) mod vendor

## mocks: generate mocks
.PHONY: mocks
mocks:
	mockery --dir=./internal --output=./mocks

## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/easypark