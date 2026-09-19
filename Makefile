ROOT_DIR := $(PWD)
BIN_DIR := $(ROOT_DIR)/build/bin
LOCAL_DB_DIR := $(ROOT_DIR)/build/dev/db
LOCAL_DB_PATH := $(LOCAL_DB_DIR)/db.sqlite3
HTMX_FILE := $(ROOT_DIR)/internal/server/web/assets/static/js/htmx.min.js
HTMX_SSE_FILE := $(ROOT_DIR)/internal/server/web/assets/static/js/htmx-ext-sse.min.js

.PHONY: build
build: build/server build/client

# todo: сhecksums
$(HTMX_FILE):
	mkdir -p $(@D)
	curl https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js --output $(@)

# todo: сhecksums
$(HTMX_SSE_FILE):
	mkdir -p $(@D)
	curl https://cdn.jsdelivr.net/npm/htmx-ext-sse@2.2.4 --output $(@)

.PHONY: build/server
build/server: $(HTMX_FILE) $(HTMX_SSE_FILE)
	go build -o $(BIN_DIR)/server ./cmd/server

.PHONY: build/client
build/client:
	go build -o $(BIN_DIR)/client ./cmd/client

.PHONY: clean
clean:
	rm -f $(BIN_DIR)/server $(BIN_DIR)/client

.PHONY: run/server
run/server: build
	$(BIN_DIR)/server -config=$(ROOT_DIR)/configs/server_local.json

.PHONY: debug/server
debug/server:
	dlv debug -l 127.0.0.1:38697 --headless $(ROOT_DIR)/cmd/server/*.go -- -config=./configs/server_local.json

.PHONY: run/client
run/client: build/client
	$(BIN_DIR)/client -config=$(ROOT_DIR)/configs/client_local.json $(if $(device),-device=$(device)) $(if $(event),-event=$(event))

.PHONY: debug/client
debug/client:
	dlv debug -l 127.0.0.1:38697 --headless $(ROOT_DIR)/cmd/client/*.go -- -config=./configs/client_local.json

.PHONY: start
start: run/server

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	golangci-lint-v2 fmt

.PHONY: lint
lint:
	golangci-lint-v2 run $(if $(fix),--fix)

.PHONY: db/create
db/create:
	rm -f $(LOCAL_DB_PATH)
	sqlite3 $(LOCAL_DB_PATH) "VACUUM; PRAGMA journal_mode=WAL; PRAGMA wal_autocheckpoint=1000; PRAGMA synchronous=NORMAL; PRAGMA busy_timeout=2000;"

.PHONY: db/create_migration
db/create_migration:
	GOOSE_MIGRATION_DIR=$(ROOT_DIR)/internal/server/pkg/db/migrations goose sqlite3 $(LOCAL_DB_PATH) create $(name) sql

.PHONY: db/populate
db/populate:
# sha256('lolkek')
	sed -e 's/__NAME__/JohnDoe/g' -e 's/__KEY__/abb45ef89186194a3e3ee700894caeb3b86ced38db1fa3ec7fd0f5e2ff6d9ec1/g' $(LOCAL_DB_DIR)/test_data.sql | sqlite3 $(LOCAL_DB_PATH)
# sha256('ololo')
	sed -e 's/__NAME__/JaneDoe/g' -e 's/__KEY__/0cb2ac8bcf600372b573bf9f807de3eb0b3ceda51c0a2045d6902412c53451f2/g' $(LOCAL_DB_DIR)/test_data.sql | sqlite3 $(LOCAL_DB_PATH)

.PHONY: gen/api
gen/api: gen/api/models gen/api/server gen/api/client

.PHONY: gen/api/models
gen/api/models:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(ROOT_DIR)/api/models.oapi-codegen.yaml $(ROOT_DIR)/api/openapi.yaml

.PHONY: gen/api/server
gen/api/server:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(ROOT_DIR)/api/server.oapi-codegen.yaml $(ROOT_DIR)/api/openapi.yaml

.PHONY: gen/api/client
gen/api/client:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(ROOT_DIR)/api/client.oapi-codegen.yaml $(ROOT_DIR)/api/openapi.yaml
