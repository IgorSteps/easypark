GO := go

CMD_DIR := ./cmd
DIST_DIR := ./dist
INTERNAL_DIR := ./internal

## build: build the project
.PHONY: build
build:
	$(GO) build -o $(DIST_DIR)/easypark $(CMD_DIR)/easypark
	$(GO) build -o $(DIST_DIR)/easypark-dbmigrate $(CMD_DIR)/easypark-dbmigrate

## setup-db: run easypark-dbmigrate to setup DB connection and migrate schemas
.PHONY: setup-db
setup-db:
	./dist/easypark-dbmigrate

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