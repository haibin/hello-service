build:
	go build ./...

install-docker:
	apt-get update
	apt-get -y install apt-transport-https ca-certificates curl gnupg lsb-release
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
	echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
	apt-get update
	DEBIAN_FRONTEND=noninteractive apt-get -y install docker-ce docker-ce-cli containerd.io

fix:
	go fmt ./...
	goimports -w .

lint:
	golangci-lint run

run:
	go run app/hello-api/main.go

test:
	go test -count=1 -coverprofile=c.out ./...

tidy:
	go mod tidy