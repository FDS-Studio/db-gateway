.DEFAULT_GOAL = run

run:
	go run ./cmd/main.go
swg:
	swag init -g ./cmd/main.go -o docs