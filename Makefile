run:
	go run server.go

testing:
	go test ./test

setup:
	sh ./setup.sh

generate:
	gqlgen generate
	sqlboiler sqlite3
	go generate ./...