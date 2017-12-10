package main

import (
	"log"
	"net"
	"strings"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/user/grpcDogParkGo/dogpark"
)

const (
	port = ":50051"
)

// server is used to implement dogpark.DogParkServer.
type server struct {
	savedDogParks []*pb.DogParkRequest
}

// CreateDogPark creates a new DogPark
func (s *server) CreateDogPark(ctx context.Context, in *pb.DogParkRequest) (*pb.DogParkResponse, error) {
	s.savedDogParks = append(s.savedDogParks, in)
	return &pb.DogParkResponse{Id: in.Id, Success: true}, nil
}

// GetDogParks returns all dogparks by given filter
func (s *server) GetDogParks(filter *pb.DogParkFilter, stream pb.DogPark_GetDogParksServer) error {
	for _, dogpark := range s.savedDogParks {
		if filter.Keyword != "" {
			if !strings.Contains(dogpark.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(dogpark); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterDogParkServer(s, &server{})
	s.Serve(lis)
}
