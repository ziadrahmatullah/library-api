ENV := release
NAME := api
SRC_APP := ./cmd/$(NAME)/$(NAME).go
BUILD_BIN := ./bin/$(NAME)
BUILD_CMD := go build -o $(BUILD_BIN) $(SRC_APP)

reload: air
	@GIN_MODE=$(ENV) air --log.main_only=true --build.cmd "$(BUILD_CMD)" --build.bin "$(BUILD_BIN)"

migrate:
	@go run ./cmd/migrate/migrate.go

seed:
	@go run ./cmd/seed/seed.go

reset: migrate seed

build:
	@$(BUILD_CMD)

run: build
	@GIN_MODE=$(ENV) $(BUILD_BIN)

test:
	@go test ./... -v

cover:
	@go test ./... --cover --coverprofile=cover.out
	@go tool cover -html=cover.out
	@rm cover.out

air:
	@command -v air > /dev/null || go install github.com/cosmtrek/air@latest

clean:
	@rm $(BUILD_BIN)