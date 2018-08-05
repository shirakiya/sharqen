.PHONY: deps migrate seed run

deps:
	dep ensure

migrate: bin/migrate.go
	go run bin/migrate.go

seed: bin/seed.go
	go run bin/seed.go

run: bin/server.go
	go run bin/server.go

runall: bin/server.go migrate seed run
