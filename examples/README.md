# Examples

Some options can be set at global level per json file.
To override few json file properties like compare/print, we can use `-G` flag along with respective override flags[`-c`, `-P` and so on]

## Global Flags
Use `-G` to enable global options. Available global flags are:
- Server address
- Proto path and proto file
- Enable/Disable TLS
- Service endpoint
- Enable/Diable stream payload
- Comparing responses
- Printing results
- Providing grpcurl flags

**NOTE**: keep an eye on global boolean flags *-t, -S, -c* and *-P*, using -G make these options `false` by default. To make better use of -G option, plan test files accordingly. Having one endpoint tests per file is recommended.

Some examples of using -G on [sayhello.json](./sayhello.json)

### Compare:
By providing `-c` option, by defalut tester compares all the test cases from the file on the expectations set. if expectations were not set then those test fails!
```shell
./grpc-tester test -j examples/sayhello.json -s localhost:8001 -e greeter.Greeter.SayHello -Gc
```
Here `-Gc` represents, override all the test cases compare property to be true.

### Print:
Similar to global compare(-c), using `-P` option overrides print property from the json file on all test cases.
```shell
./grpc-tester test -j examples/sayhello.json -s localhost:8001 -e greeter.Greeter.SayHello -GP
```
Since `-c` is not provided, comparing responses with expectations is ignored.

### grpcurl flags:
To provide additional grpcurl flags for all the test cases, use `-g`.
```shell
./grpc-tester test -GPg '-H "Authorization: eyusicf"' -j examples/greeter-test.json
```
