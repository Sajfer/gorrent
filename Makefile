ld:
	go build -o bin/gorrent main.go

test:
	go fmt ./...
	go vet ./...
	go test -race -v ./...

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/gorrent main.go
