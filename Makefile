# Variables
DOCKER_COMPOSE_FILE=docker-compose.yaml
SERVICE_NAME=pdf-generator-app
TEST_SCRIPT=run_parallel_curl
NUM_REQUESTS=100
MAX_CONCURRENT=50

# Default target
all: build run

# Build the Docker image
build:
	@echo "Building Docker image..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Start the services
run:
	@echo "Starting services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop the services
stop:
	@echo "Stopping services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Test the API with parallel curl requests
test:
	@echo "Testing API with $(NUM_REQUESTS) parallel requests..."
	./$(TEST_SCRIPT) $(NUM_REQUESTS) $(MAX_CONCURRENT)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf output_pdfs

# View logs for the PDF generator service
logs:
	@echo "Viewing logs for $(SERVICE_NAME)..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f $(SERVICE_NAME)

# Help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build       Build the Docker image"
	@echo "  run         Start the services"
	@echo "  stop        Stop the services"
	@echo "  test        Test the API with parallel curl requests"
	@echo "  clean       Clean up generated files"
	@echo "  logs        View logs for the PDF generator service"
	@echo "  help        Show this help message"

.PHONY: all build run stop test clean logs help