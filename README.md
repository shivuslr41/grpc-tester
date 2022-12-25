# grpc-tester

`grpc-tester` is a simple command-line tool cum library for testing gRPC services.

This tool leverage [grpcurl](https://github.com/fullstorydev/grpcurl) to create dynamic requests and [jq](https://github.com/stedolan/jq) for comparing/validating responses.

NOTE: currently doesn't work on windows cmd/powershell.

DOCUMENTATION is work in progress, have a look at `examples/` to get overview of tool usage!

### Install From Source
```shell
go install github.com/shivuslr41/grpc-tester/cmd/grpc-tester@v0.1.0-beta
```

## Usage
```
$ grpc-tester --help

FORMAT:
        ./tester [COMMAND] [FLAGS]

EXAMPLE:
        ./tester list --server mygrpcserver:443 --tls

COMMANDS:
        gen             generates sample json.
        list            list services and methods.
        run             run requests provided in json.
        test            test responses againt expectation set.

Usage of ./grpc-tester:
  -d, --data string            request in json format - '{"name":"Bob"}'
  -e, --endpoint string        service and method to call
  -g, --grpcurl-flags string   pass additional grpcurl flags - '-H "Authorization: <TOKEN>"'
  -h, --help                   shows tool usage
  -j, --json string            json file containing test scopes
  -f, --proto-file string      proto file
  -p, --proto-path string      proto path, if server doesn't support grpc reflection
  -s, --server string          gRPC server address
  -m, --stream-payload         send multiple messages to server
  -t, --tls                    use tls connection
```

## Examples
### gen
This generates json format file, add scopes and pass it to grpc-tester run/test commands.
```shell
grpc-tester gen 
```
