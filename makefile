# Variables
APP_NAME := mvrp-api
APP_CMD := app
DOCKER_IMAGE := $(APP_NAME):latest
DOCKERFILE := dockerfile

# Docker build arguments
PORT := 6900
DB_PORT := 5432
DB_HOST := db

###########################################################################################
# APP
###########################################################################################

# Run application
.PHONY: run
run:
	@echo "Starting application..."
	go run ./cmd/app/main.go

# Run application in dev mode
.PHONY: run-dev
run-dev:
	@echo "Starting application in development mode..."
	cd cmd/app && air

###########################################################################################
# DOCKER
###########################################################################################
# Default target
.PHONY: all
all: build docker

# Build the Go binary
.PHONY: build
build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME) ./cmd/$(APP_CMD)/

# Build the Docker image
.PHONY: docker
docker: build
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	@echo "Running the Docker container..."
	docker run --name $(APP_NAME)-container -p $(PORT):$(PORT) --env PORT=$(PORT) --env DB_HOST=$(DB_HOST) --env DB_PORT=$(DB_PORT) $(DOCKER_IMAGE)

# Stop and remove the Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping and removing the Docker container..."
	docker stop $(APP_NAME)-container || true
	docker rm $(APP_NAME)-container || true

# Clean up the binary
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

###########################################################################################
# CODEGEN
###########################################################################################

# Run codegen
.PHONY: gen
gen:
	@echo "Generating code..."
	go run ./cmd/gen/main.go

###########################################################################################
# SEED
###########################################################################################

# Seed the database
.PHONY: seed
seed:
	@echo "Seeding the database..."
	go run ./cmd/seed/.