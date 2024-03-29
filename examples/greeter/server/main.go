package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	greeterpb "github.com/shivuslr41/grpc-tester/examples/greeter/proto/greeter"
)

// server pb
type server struct {
	greeterpb.UnimplementedGreeterServer
}

// SayHello greets the client with provided name
func (s *server) SayHello(ctx context.Context, in *greeterpb.HelloRequest) (*greeterpb.HelloReply, error) {
	log.Println("SayHello called")
	if in.GetName() == "" {
		return nil, status.Errorf(codes.NotFound, "%s", "name not found")
	}
	return &greeterpb.HelloReply{Message: "Hello " + in.GetName() + "!"}, nil
}

// SayHelloStream greets the client in streams
func (s *server) SayHelloStream(in greeterpb.Greeter_SayHelloStreamServer) error {
	log.Println("SayHelloStream called")
	for {
		select {
		case <-in.Context().Done():
			log.Println("context canceled from client")
			return nil
		default:
		}
		res, err := in.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("received EOF from client")
				return nil
			}
			return err
		}
		out := &greeterpb.HelloReply{}
		if res.GetName() == "" {
			out.Message = "invalid name"
		} else {
			out.Message = "Hello " + res.GetName() + "!"
		}
		err = in.Send(out)
		if err != nil {
			return err
		}
	}
}

var (
	silent bool
)

func init() {
	flag.BoolVar(&silent, "silent", false, "enable/disable printing logs")
	flag.Parse()
}

// start demo server
func main() {

	if silent {
		// Disable log output
		log.SetOutput(io.Discard)
	}

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8333")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	greeterpb.RegisterGreeterServer(s, &server{})
	// register reflection endpoint
	reflection.Register(s)
	// Register the Health service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	// Serve gRPC Server
	log.Println("Serving gRPC on", lis.Addr().String())
	log.Fatal(s.Serve(lis))
}
