# server 
server:
	@echo "Starting server"
	@go run main.go

# Create a new postgres migration. make create-migration-postgres name="migration name"
create-migration-postgres:
	@echo "Creating migration"
	@goose -dir infra/postgres/migrations create $$name go

# Up the migrations
migration-up-postgres:
	@echo "Up the migrations"
	@go run infra/goose/migrations.go up

sol:
	solc --optimize --abi ./contracts/MySmartContract.sol -o build
	solc --optimize --bin ./contracts/MySmartContract.sol -o build
	abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go