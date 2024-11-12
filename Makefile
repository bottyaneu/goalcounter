build:
	@go build -o ./bin/goalcounter

run: build
	@./bin/goalcounter --listenAddr :3002
