init:
	npm install
	
build:
	@echo "Building frontend..."
	npm run build

build-dev:
	@echo "Building frontend in development mode..."
	npm run start

run: build
	docker-compose -f ../docker-compose.yml up -d frontend

build-run: build
	docker-compose -f ../docker-compose.yml up -d --build frontend

build-platform:
	@echo "Building frontend for specific platform..."
	docker build -t frontend:latest .
