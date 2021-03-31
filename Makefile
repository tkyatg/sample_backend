up:
	docker-compose up -d
down:
	docker-compose down
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go
gen-mock:
	mockgen -source shared/env.go -destination shared/env_mock.go -package shared
	# mockgen -source auth/commands/authenticationCommand/usecase.go -destination auth/commands/authenticationCommand/usecase_mock.go -package authenticationcommand
	mockgen -source services/user/queries/userQuery/usecase.go -destination services/user/queries/userQuery/usecase_mock.go -package userquery
test:
	go test -race -covermode=atomic ./... -test.v