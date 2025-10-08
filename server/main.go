package main

import (
	"log"
	"net"

	"nebula/orchestrator"
	pb "nebula/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Unable to load .env file.")
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWorkflowServiceServer(grpcServer, &orchestrator.WorkflowServer{})

	log.Println("🚀 Nebula Orchestrator gRPC server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
