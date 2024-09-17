APP_NAME=hdflex-api
ROOT_PATH=/srv/hdflex.tv
DOCKER_EXE=docker compose
DOCKER_NETWORK=hdflex-network

.PHONY: docker

docker: check-network create-env
	@echo "Starting the Docker Compose stack..."
	$(DOCKER_EXE) up -d

build-api:
	@cd "$(ROOT_PATH)/cmd/api" && go build -buildvcs=false -o "$(APP_NAME)"
	@chmod +x "$(ROOT_PATH)/cmd/api/$(APP_NAME)"
	@mv "$(ROOT_PATH)/cmd/api/$(APP_NAME)" "$(ROOT_PATH)/tmp/$(APP_NAME)"

check-network:
	@echo "Checking if the network exists..."
	@docker network inspect $(DOCKER_NETWORK) > /dev/null 2>&1 || (echo "Network does not exist. Please create it." && docker network create $(DOCKER_NETWORK))
	@echo "Network exists."

create-env:
	@if [ ! -f .env ]; then echo "Creating .env file..."; cp .env.example .env; fi

database:
	@make -C internal/database $(action)

balancer:
	@rm -rf storage/public
	@go run cmd/balancer/main.go

log:
	@mkdir -p storage/logs