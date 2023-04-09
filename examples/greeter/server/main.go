package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	if in.GetName() == "" {
		return nil, status.Errorf(codes.NotFound, "%s", "name not found")
	}
	return &greeterpb.HelloReply{Message: "Hello " + in.GetName() + "!"}, nil
}

// SayHelloStream greets the client in streams
func (s *server) SayHelloStream(in greeterpb.Greeter_SayHelloStreamServer) error {
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

// start demo server
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	greeterpb.RegisterGreeterServer(s, &server{})
	// register reflection endpoint
	reflection.Register(s)
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8001")
	log.Fatal(s.Serve(lis))
}
