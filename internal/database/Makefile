DB_USER=root
DB_PASSWORD=root
DB_HOST=mysql
DB_PORT=3306
DB_NAME=hdflex

.PHONY: up

migration:
	@echo "Creating the migrations..."
	@migrate create -ext sql -dir ./migrations -seq $(name)

fresh: down up

up:
	@echo "Running the migrations..."
	@./.bin/migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

down:
	@echo "Rolling back the migrations..."
	@./.bin/migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down