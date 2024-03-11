run:
	@templ generate
	@go run cmd/main.go

build:
	@go build -o ./bin/recallme ./cmd
