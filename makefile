tools:
	go build -o ./tool ./cmd/tools/  && ./tool

gen:
	go generate ./...

run:
	go run ./cmd/api/main.go

docs:
	swag init --generalInfo internal/api/api.go --dir ./ --parseInternal
