build:
	go build ./...

fmt:
	go fmt ./...

run:
	go run app/hello-api/main.go

test:
	go test ./... -count=1
	staticcheck ./...

tidy:
	go mod tidy