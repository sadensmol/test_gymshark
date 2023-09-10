.PHONY up:
up:
	go run github.com/cosmtrek/air@latest

.PHONY test:
test:
	go test -v ./...