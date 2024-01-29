APP_NAME=root
ROOT_PATH=/app
DOCKER_EXE=docker compose
DOCKER_NETWORK=chat-network

.PHONY: docker

docker: check-network create-env
	@echo "Starting the Docker Compose stack..."
	$(DOCKER_EXE) up -d

build:
	@cd "$(ROOT_PATH)/cmd/$(APP_NAME)" && go build -buildvcs=false -o "$(APP_NAME)"
	@chmod +x "$(ROOT_PATH)/cmd/$(APP_NAME)/$(APP_NAME)"
	@mv "$(ROOT_PATH)/cmd/$(APP_NAME)/$(APP_NAME)" "$(ROOT_PATH)/tmp/$(APP_NAME)"
	@make ws

ws:
	@echo "Building ws..."
	@cd "$(ROOT_PATH)/cmd/ws" && go build -buildvcs=false -o "ws"
	@chmod +x "$(ROOT_PATH)/cmd/ws/ws"
	@mv "$(ROOT_PATH)/cmd/ws/ws" "$(ROOT_PATH)/tmp/ws"
	@echo "ws built."
	@bash -c "$(ROOT_PATH)/tmp/ws" 2>&1 &

check-network:
	@echo "Checking if the network exists..."
	@docker network inspect $(DOCKER_NETWORK) > /dev/null 2>&1 || (echo "Network does not exist. Please create it." && docker network create $(DOCKER_NETWORK))
	@echo "Network exists."

create-env:
	@if [ ! -f .env ]; then echo "Creating .env file..."; cp .env.example .env; fi


