# Variables
BINARY_NAME = chatchk
MAIN_MODULE = ./src/chatchk
MODULES = ingest knowledge prompts admin
REGISTRY = ddonahuex
IMAGE_NAME = chatchk
TAG ?= 1.0.0
IMAGE = $(REGISTRY)/$(IMAGE_NAME):$(TAG)

# Default target
all: build
prod: build docker-prod

# Build the chatchk executable
build:
	@echo "Tidying modules..."
	@for dir in $(MAIN_MODULE) $(MODULES); do \
		cd $$dir && go mod tidy && cd ..; \
	done
	@echo "Building $(BINARY_NAME)..."
	@cd $(MAIN_MODULE) && go build -o ../$(BINARY_NAME)

# Run tests for all modules
test:
	@for dir in $(MAIN_MODULE) $(MODULES); do \
		echo "Testing $$dir..."; \
		cd $$dir && go test ./... && cd ..; \
	done

# Clean up
clean:
	@echo "Cleaning..."
	@rm -f $(MAIN_MODULE)/$(BINARY_NAME)
	@for dir in $(MAIN_MODULE) $(MODULES); do \
		cd $$dir && go clean && cd ..; \
	done

# Build Docker build & push
docker-prod: docker-sbom docker-push

# Build Docker image
docker-build:
	@echo "Building Docker image ..."
	@docker build -t $(IMAGE) .

# Build & run Docker container
docker-run: docker-build
	@echo "Running Docker container ..."
	@docker run -it --rm -p 8080:8080 $(IMAGE)

docker-push:
	@echo "Running Docker push ..."
	@docker push $(IMAGE)

docker-sbom: docker-build
	@echo "Generating SBOM (CycloneDX) ..."
	@syft docker:$(IMAGE) -o cyclonedx-json | jq . > temp.json && mv temp.json $(IMAGE_NAME):$(TAG)-bom.json

help:
	@echo "Make Targets:"
	@echo "  build\t\tExecutes a Go build for the $(IMAGE_NAME) executable"
	@echo "  clean\t\tExecutes a Go clean for all modules"
	@echo "  docker-build\tDocker build for $(REGISTRY)/$(IMAGE_NAME) Docker image"
	@echo "  docker-prod\tExecutes docker-build and docker-push make targets"
	@echo "  docker-push\tDocker push for of $(REGISTRY)/$(IMAGE_NAME) to the $(REGISTRY) Docker Hub namespace"
	@echo "  docker-run\tExecutes docker-build then issues a docker run of the $(REGISTRY)/$(IMAGE_NAME) image"
	@echo "  docker-sbom\tExecutes docker-build then generates CycloneDX formatted SBOM using syft"
	@echo "  help\t\tPrint this help menu"
	@echo "  prod\t\tExecutes build and docker-prod make targets"
	@echo "  test\t\tExecutes a Go test for all modules"

	
.PHONY: all prod build test clean docker-all docker-build docker-run docker-push help chatchk