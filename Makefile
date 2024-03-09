GO := go

CMD_DIR := ./cmd
DIST_DIR := ./dist

## build: build the project
.PHONY: build
build:
	$(GO) build -o $(DIST_DIR)/easypark $(CMD_DIR)/easypark

## run: run easypark 
.PHONY: run
run:
	./dist/easypark