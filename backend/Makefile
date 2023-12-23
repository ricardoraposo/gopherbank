build:
	@echo "Building..."
	@go build -o bin/gopherbank cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

gen:
	@echo "Generating..."
	@go mod tidy
	@go generate ./ent/

migrate: gen
	@echo "Migrating..."
	@go run cmd/migration/main.go

seed:
	@echo "Seeding..."
	@go run cmd/seed/main.go
