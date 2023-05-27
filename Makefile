# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = build/zus_server
MAIN_PACKAGE = ./cmd/zus-server

MIGRATE_URL = https://github.com/golang-migrate/migrate/releases/download/v4.4.0/migrate.linux-amd64.tar.gz
MIGRATE_FILE = migrate.linux-amd64.tar.gz
MIGRATE_DIR = ./migrate.linux-amd64
MIGRATE_COMMAND = $(MIGRATE_DIR)/migrate.linux-amd64

DB_USER = your_db_user
DB_PASS = your_db_password
DB_ADDRESS = your_db_address

.PHONY: all deps clean build test

all: deps build

# Download and install dependencies
deps:
	$(GOGET) ./...

# Clean build files and binaries
clean:
	$(GOCLEAN) -testcache
	rm -rf build
	mkdir -p build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run tests
test:
	$(GOTEST) ./... -p 1 -run TestDatabase
	$(GOTEST) ./... -p 1 -run "[^(TestDatabase)]"

# Perform database creation and migration
setupdb:
	echo "SETUP DB"
	psql -U postgres -c "CREATE USER $(DB_USER);"
	psql -U postgres -c "ALTER USER $(DB_USER) WITH ENCRYPTED PASSWORD '$(DB_PASS)';"
	psql -U postgres -c "CREATE DATABASE \"zus\";"
	psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE zus TO $(DB_USER);"
	psql -U $(DB_USER) -d zus -c "CREATE EXTENSION \"uuid-ossp\";"

	echo "MIGRATE DB"
	wget $(MIGRATE_URL) -O $(MIGRATE_FILE)
	tar -xzf $(MIGRATE_FILE)
	$(MIGRATE_COMMAND) -source file://database/migrations -database postgres://$(DB_USER):$(DB_PASS)@tcp($(DB_ADDRESS):3306)/zus goto 0202
	rm -rf $(MIGRATE_FILE) $(MIGRATE_DIR)