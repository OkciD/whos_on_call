ROOT_PATH := $(PWD)
BIN_PATH := $(ROOT_PATH)/build/bin
LOCAL_DB_PATH := $(ROOT_PATH)/build/dev/db/db.sqlite3

.PHONY: build
build: build/bin/whos_on_call

build/bin/whos_on_call:
	go build -o $(@) ./cmd/whos_on_call

.PHONY: clean
clean:
	rm -f $(BIN_PATH)/*

.PHONY: run
run: clean build
	$(BIN_PATH)/whos_on_call -config=$(ROOT_PATH)/configs/local.json

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
