# grpc-tester

`grpc-tester` is a simple go-lib and command-line tool for testing gRPC services.

Tool leverages [grpcurl](https://github.com/fullstorydev/grpcurl) to execute gRPC requests and [jq](https://github.com/stedolan/jq) for filtering and validating responses.

NOTE: Doesn't work on windows cmd/powershell.

### Prerequisites
- [grpcurl](https://github.com/fullstorydev/grpcurl#installation)
- [jq](https://stedolan.github.io/jq/download/)

### Install From Source
```shell
go install github.com/shivuslr41/grpc-tester/cmd/grpc-tester@v0.1.0-beta
```

## Usage
```
$ grpc-tester --help

FORMAT:
        ./grpc-tester [COMMAND] [FLAGS]

EXAMPLE:
        ./grpc-tester list --server mygrpcserver:443 --tls

COMMANDS:
        gen             generates sample json.
        list            list services and methods.
        run             run executes requests provided in json or via -d flag.
        test            test responses againt expectations set.

Usage of ./grpc-tester:
  -c, --compare                test/compare responses
  -d, --data string            request in json format - '{"name":"Ramesh"}'
  -e, --endpoint string        service and method to call
  -G, --global                 consider global flags for run/test commands
  -g, --grpcurl-flags string   pass additional grpcurl flags - '-H "Authorization: <TOKEN>"'
  -h, --help                   shows tool usage
  -j, --json string            json file containing test scopes
  -P, --print                  prints result
  -f, --proto-file string      proto file
  -p, --proto-path string      proto path, if server doesn't support grpc reflection
  -s, --server string          gRPC server address
  -S, --stream-payload         send multiple messages to server
  -t, --tls                    use tls connection
```

## Examples
### gen
generates json file for the reference.
```shell
grpc-tester gen 
```

### list
lists all avaiable services and methods on a given server.
```shell
grpc-tester list -s localhost:8001
```
use `-t`/`--tls` for secure server connections.
```shell
grpc-tester list -ts myserver:443
```

### run
executes requests provided via `-d`/`--data` flag. multiple requests can be executed by providing them in json file's `tests` property (set `compare` as false and `print` as true for each test cases to make them just run cases).
```shell
grpc-tester run -s localhost:8001 -d '{"name":"Ramesh"}'
```
streaming payload can be provided by using `-S`/`--stream-payload` flag and enclosing multiple stream requests in `[]`.
```shell
grpc-tester run -s localhost:8001 -d '[{"name":"Ramesh"},{"name":"Suresh"}]' -S
```

### test
validate responses against set expectations in the json file after filtering the responses by the given `jq` queires in the json file via `jqq` property. refer `examples/greeter-test.json` for more details.
```shell
grpc-tester test -j examples/greeter-test.json
```

## JSON format
```json
[
    {
        "server": "",
        "proto_path": "",
        "proto_file": "",
        "tls": false,
        "endpoint": "",
        "stream_payload": false,
        "tests": [
            {
                "id": "",
                "description": "",
                "requests": [
                    {}
                ],
                "jqq": [
                    ""
                ],
                "compare": false,
                "expectations": [
                    {}
                ],
                "skip": false,
                "print": false,
                "grpcurl_flags": "",
                "ignore_order": false
            }
        ]
    }
]
```
TODO: add json details.
