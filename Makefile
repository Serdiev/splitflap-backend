build:
	go build -o bin/main cmd/splitflap/main.go

run:
	go run cmd/splitflap/main.go

sand:
	go run cmd/sandbox/main.go

runn: 
	go build -o app.exe cmd/splitflap/main.go && app.exe


test:
	go test ./...