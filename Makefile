SERVER_BIN=server
MIGRATOR_BIN=migrator

SERVER_CMD=./cmd/app
MIGRATOR_CMD=./cmd/migrator
MIGRATIONS_DIR=./migrations
CONFIG_DIR=./config
SCRIPTS_DIR=./scripts

build:
	go build -o $(SERVER_BIN) "$(SERVER_CMD)"
	go build -o $(MIGRATOR_BIN) "$(MIGRATOR_CMD)"

migrate-up: build
	@echo "Running UP migrations..."
	./$(MIGRATOR_BIN) --config="$(CONFIG_DIR)/local.yaml" --dir="$(MIGRATIONS_DIR)" --direction=up

migrate-down: build
	@echo "Running DOWN migrations..."
	./$(MIGRATOR_BIN) --config="$(CONFIG_DIR)/local.yaml" --dir="$(MIGRATIONS_DIR)" --direction=down

add-timestamp:
	chmod +x "$(SCRIPTS_DIR)/add-timestamp.sh"
	"$(SCRIPTS_DIR)/add-timestamp.sh"

lint:
	golangci-lint run

test:
	go test ./...