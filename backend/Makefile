.PHONY: deps migrate seed run

deps:
	dep ensure

migrate: bin/migrate.go
	go run bin/migrate.go

seed: bin/seed.go
	go run bin/seed.go

run: bin/server.go migrate
	go run bin/server.go
