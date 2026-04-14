.PHONY: dev build run gen tidy clean install-tools

APP = lagarigo

install-tools:
	go install github.com/a-h/templ/cmd/templ@latest

gen:
	templ generate

tidy:
	go mod tidy

dev: gen
	go run ./cmd/server

build: gen
	go build -o bin/$(APP) ./cmd/server

run: build
	./bin/$(APP)

clean:
	rm -rf bin/ *.db
	find . -name "*_templ.go" -type f -delete
