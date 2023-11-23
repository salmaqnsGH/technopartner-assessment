createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5455/belajar_golang_restful_api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5455/belajar_golang_restful_api?sslmode=disable" -verbose down


.PHONY: createmigrate migrateup migratedown