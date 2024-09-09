build:
	GOOS=linux GOARCH=arm64 go build -o bin/golangBuild cmd/splitflap/main.go

ex:
	nohup ./bin/golangBuild &

id:
	ps -e | grep "golangBuild" | awk '{print $1}'

# "GOOS=linux GOARCH=arm64" needed to run on raspberry pi 4. Can skip otherwise.
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
