# grpc-tester
`grpc-tester` is a simple go-lib and command-line tool for testing gRPC services.

## Introduction
The gRPC Tester is a tool for testing gRPC APIs. It provides an easy and efficient way for developers to verify the functionality of their gRPC APIs. With a simple command-line interface, users can run tests and check the results of their API calls. Whether you're a seasoned gRPC developer or just getting started, the gRPC Tester makes it easy to test and debug your APIs, ensuring that they work as expected before deployment.

Tool leverages [grpcurl](https://github.com/fullstorydev/grpcurl) to execute gRPC requests and [jq](https://github.com/stedolan/jq) for filtering, validating responses and dynamic data injection from one request to another.

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
        "skip": false,
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
                "ignore_order": false,
                "replace": null,
                "replace_from": null,
                "extract": null,
                "extract_to": null
            }
        ]
    }
]
```
- **server**: add grpc server addr and port.
- **endpoint**: provide the service and method you want to request/test.
- **proto_path**: if the server doesn’t support gRPC reflection then proto files should be used.
- **proto_file**: use the proto file from the given proto path.
- **tls**: *true* if server supports tls connections (for passing certs use *grpcurl-flags*).
- **stream_payload**: *true* for streaming requests (client-streaming and bi-directional stream).
- **skip**: set *true* if all the tests for a given method/endpoint need to be skipped.
- **tests**: include all required tests/requests and expectations.
  - **id**: unique id just for separating tests from each other.
  - **description**: provide test description.
  - **print**: *true* for printing responses/errors to stdout/stderr.
  - **requests**: provide a JSON object(s) that contains the request.
  - **grpcurl-flags**: allow us to provide additional supported `grpcurl` flags (refer [*grpcurl*](https://github.com/fullstorydev/grpcurl#usage) docs and refer [*examples/README.md*](./examples/README.md) for more details).
  - **compare**: *true* to compare all responses with set *expectations*.
  - **jqq**: *jq query* can be used to select particular nested JSON object/field from responses. default selects the whole result to compare against the set expectation.</br>
  *example* - ".data" selects the nested data field from the result to compare against set *expectations* expectation.</br>
  (refer [*examples/README.md*](./examples/README.md) for more details)
  - **ignore_order**: if set to *true* ignores order while comparing responses against *expectations*.
  - **expectations**: set expectations that need to be validated against responses. (mandatory if *compare* is true)
  - **comments**: add any comments/reasons if necessary on test failures.
  - **skip**: *true* if the test needs to be skipped.
  - **extarct**: part of response can be extarcted and stored in temporary json *variables.json* so that this extarcted data can be used to pass as a part of another endpoint/method's request.
  - **extract_to**: variable name (feild name of *variables.json*) in which extarcted data to be stored, later this can be used in *replace_from*.
  - **replace**: can be used to replace *requests* nested object / feild data from *extract*ed feild/object data from previous/another endpoint/methods result.
  - **replace_from**: select the variable name (feild name of *variables.json*) from which data to be considered for replace.</br>
  (refer [*examples/README.md*](./examples/README.md) for more details)

## Similar projects
> grpc-tester was inspired by testing tool we built at flexera for testing internal grpc APIs. The codebase was written in bash script using same dependencies `grpcurl` and `jq`. I've re-written same using golang with few more features for better outreach and to ship this as an installable package. In this section I just want to mention other good projects.
- [grpc-testing](https://github.com/ryoya-fujimoto/grpc-testing)
- [karate-grpc](https://github.com/pecker-io/karate-grpc)

## Support the project!
If you like the project, consider giving a ⭐️! It draws more attention to the project, which helps us improve it even faster.
