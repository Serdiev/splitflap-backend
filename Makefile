build:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/test cmd/splitflap/main.go
# # CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o bin/main cmd/splitflap/main.go

run:
	GOOS=linux GOARCH=arm64 go run cmd/splitflap/main.go

sand:
	go run cmd/sandbox/main.go

runn: 
	go build -o app.exe cmd/splitflap/main.go && app.exe


test:
	go test ./...


cert:
	sudo certbot certonly --standalone
