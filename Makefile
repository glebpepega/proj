up:
	migrate -path migrations -database "postgresql://gleb:1510@localhost:5432/proj?sslmode=disable" -verbose up

down:
	migrate -path db/migration/ -database "postgresql://gleb:1510@localhost:5432/proj?sslmode=disable" -verbose down