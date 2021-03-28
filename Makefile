up:
	docker-compose up -d
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go