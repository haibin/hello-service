build:
	go build ./...

lint:
	golangci-lint run

fix:
	go fmt ./...
	goimports -w .

run:
	go run app/hello-api/main.go

test:
	go test ./... -count=1

tidy:
	go mod tidy