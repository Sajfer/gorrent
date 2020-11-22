.PHONY: build
build:
	go build -o bin/gorrent main.go

.PHONY: test
test:
	go fmt ./...
	go vet ./...
	go test -race -cover -v ./...

.PHONY: run
run:
	go run main.go

.PHONY: compile
compile:
	GOOS=linux GOARCH=386 go build -o bin/gorrent main.go

.PHONY: clean
clean:
	rm -rf bin/
