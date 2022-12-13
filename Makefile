fmt:
	@go fmt ./...

build:
	@go build -o grpctester cmd/*

test:
	./grpctester test -j examples/greeter-test.json
