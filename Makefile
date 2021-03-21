.PHONY: bin/api

all: .PHONY

bin/api:
	go build -o ./bin/api ./cmd/api

clean:
	rm bin/*

install:
	mv bin/api /var/www/flagship
	sudo service flagship restart

lint:
	gopls ./...

test:
	go test ./...

start:
	go run ./cmd/api/main.go

migrate:
	go run migrations/*.go migrate
	
build:
	docker run --rm -v "$(CURDIR)":/usr/src/myapp -w /usr/src/myapp golang:1.15 make
