fmt:
	@go fmt ./...

lint:
	@go install golang.org/x/lint/golint@latest && golint ./...

build:
	@go build -o grpc-tester cmd/grpc-tester/*

test:
	./grpc-tester test -j examples/greeter-test.json
