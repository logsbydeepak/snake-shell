.PHONY: tidy, build, clean, log, dev, start

dev:
	go run main.go

start:
	./build/snake-shell

tidy:
	go fmt ./...
	go mod tidy -v

build:
	go build -o build/snake-shell main.go

clean:
	rm -rf build

log:
	tail -f app.log

