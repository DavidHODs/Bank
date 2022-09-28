postgresUser = $(shell cat ${HOME}/.postgresUser)
postgresPassword = $(shell cat ${HOME}/.postgresPassword)
target1:=${postgresUser}
target2= ${postgresPassword}

schemaInit:
	migrate create -ext sql -dir db/migration/ -seq init_schema

migrateUp: 
	migrate -path db/migration/ -database "postgresql://${target1}:${target2}@localhost:5432/bank" -verbose up

migrateDown: 
	migrate -path db/migration/ -database "postgresql://${target1}:${target2}@localhost:5432/bank" -verbose down

sqlc:
	sqlc generate

PHONY: schemaInit migrateUp migrateDown sqlc