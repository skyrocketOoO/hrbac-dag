run:
	go run cmd/main.go

clear-db:
	rm ./gorm.db

test-run: clear-db run
