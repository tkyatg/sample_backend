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
	# mockgen -source services/auth/commands/authenticationCommand/usecase.go -destination services/auth/commands/authenticationCommand/usecase_mock.go -package authenticationcommand
	mockgen -source services/user/queries/userQuery/usecase.go -destination services/user/queries/userQuery/usecase_mock.go -package userquery
	mockgen -source services/user/commands/userCommand/usecase.go -destination services/user/commands/userCommand/usecase_mock.go -package usercommand
	mockgen -source services/user/domain/userRepository.go -destination services/user/domain/userRepository_mock.go -package domain
test:
	go test -race -covermode=atomic ./... -test.v