migrate:
	go run ./cmd/migrate/migrate.go

seed:
	go run ./cmd/seed/seed.go

api:
	go run ./cmd/api/api.go

test:
	go test ./... -v

cover:
	go test ./... --cover --coverprofile=cover.out
	go tool cover -html=cover.out
	rm cover.out