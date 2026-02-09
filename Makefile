ROOT_DIR := $(PWD)
BIN_DIR := $(ROOT_DIR)/build/bin
LOCAL_DB_DIR := $(ROOT_DIR)/build/dev/db
LOCAL_DB_PATH := $(LOCAL_DB_DIR)/db.sqlite3

.PHONY: build
build: build/bin/server

build/bin/server:
	go build -o $(@) ./cmd/server

.PHONY: clean
clean:
	rm -f $(BIN_DIR)/*

.PHONY: run
run: clean build
	$(BIN_DIR)/server -config=$(ROOT_DIR)/configs/server_local.json

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
	rm -f $(LOCAL_DB_PATH)
	sqlite3 $(LOCAL_DB_PATH) "VACUUM;"

.PHONY: db/create_migration
db/create_migration:
	GOOSE_MIGRATION_DIR=$(ROOT_DIR)/internal/pkg/db/migrations goose sqlite3 $(LOCAL_DB_PATH) create $(name) sql

.PHONY: db/populate
db/populate:
	sqlite3 $(LOCAL_DB_PATH) < $(LOCAL_DB_DIR)/test_data.sql
