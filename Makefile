BINARY := biathlon-competitions

.PHONY: run

run:
	go run cmd/biathlon-competitions/main.go

build:
	go build -o $(BINARY) cmd/biathlon-competitions/main.go

run-bin: build
	./biathlon-competitions

clean:
	rm -rf $(BINARY)