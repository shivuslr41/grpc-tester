fmt:
	@go fmt ./...

lint:
	@go install golang.org/x/lint/golint@latest && golint ./...

build:
	@go build -o grpc-tester cmd/grpc-tester/*

gen:
	@./grpc-tester gen

run:
	./grpc-tester run -d '{"name":"Ramesh"}' -s localhost:8001 -e greeter.Greeter.SayHello

test:
	./grpc-tester test -j examples/greeter-test.json
