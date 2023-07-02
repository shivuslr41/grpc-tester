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

start-greeter:
	go build -o greeter examples/greeter/server/*
	./greeter --silent &

wait-for-greeter:
	@echo "Waiting for greeter to be ready..."
	@timeout 30 bash -c 'until [ $$(grpcurl -plaintext localhost:8333 grpc.health.v1.Health/Check | jq -r ".status") = "SERVING" ]; do sleep 1; done' || (echo "Server startup timeout" && exit 1)

ci: build start-greeter wait-for-greeter
	./grpc-tester test -j examples/greeter-test.json
	killall greeter
