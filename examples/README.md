# Examples

TODO: add details...

G - use global flags. global flags overrides json properties if provided.

c - compares

./grpc-tester test -j examples/sayhello.json -s localhost:8001 -e greeter.Greeter.SayHello -Gc

P - prints

./grpc-tester test -j examples/sayhello.json -s localhost:8001 -e greeter.Greeter.SayHello -GP

just overwrites grpcurl flag

./grpc-tester test -GSPg '-H "Authorization: eyusicf"' -j examples/greeter-test.json