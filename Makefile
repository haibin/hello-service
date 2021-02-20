build:
	go build ./...

fmt:
	go fmt ./...
	goimports -w .

run:
	go run app/hello-api/main.go

test:
	go test ./... -count=1
	staticcheck ./...

tidy:
	go mod tidy