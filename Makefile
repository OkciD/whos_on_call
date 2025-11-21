ROOT_PATH := $(PWD)
BIN_PATH := $(ROOT_PATH)/build/bin

.PHONY: build
build: build/bin/whos_on_call

build/bin/whos_on_call:
	go build -o $(@) ./cmd/whos_on_call

.PHONY: clean
clean:
	rm -f $(BIN_PATH)/*

.PHONY: run
run: clean build
	$(BIN_PATH)/whos_on_call
