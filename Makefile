ROOT_DIR := $(PWD)
BIN_DIR := $(ROOT_DIR)/build/bin
LOCAL_DB_DIR := $(ROOT_DIR)/build/dev/db
LOCAL_DB_PATH := $(LOCAL_DB_DIR)/db.sqlite3

.PHONY: build
build: build/bin/whos_on_call

build/bin/whos_on_call:
	go build -o $(@) ./cmd/whos_on_call

.PHONY: clean
clean:
	rm -f $(BIN_DIR)/*

.PHONY: run
run: clean build
	$(BIN_DIR)/whos_on_call -config=$(ROOT_DIR)/configs/local.json

.PHONY: start
start: run

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: db/create
db/create:
	sqlite3 $(LOCAL_DB_PATH) "VACUUM;"

.PHONY: db/create_migration
db/create_migration:
	GOOSE_MIGRATION_DIR=$(ROOT_DIR)/internal/pkg/db/migrations goose sqlite3 $(LOCAL_DB_PATH) create $(name) sql

.PHONY: db/populate
db/populate:
	sqlite3 $(LOCAL_DB_PATH) < $(LOCAL_DB_DIR)/test_data.sql
