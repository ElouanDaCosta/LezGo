.PHONY: all clean

binary   = lezgo
version  = 0.1.0
build	   = $(shell git rev-parse HEAD)

all:
	go build -o ./bin/$(binary) ./cmd/main.go

test:
	go test ./... -cover -coverprofile c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm -rf $(binary) c.out coverage.html