package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/user/dogParkGrpc/dogpark"
)

const (
	address = "localhost:50051"
)

// createDogPark calls the RPC method CreateDogPark of DogParkServer
func createDogPark(client pb.DogParkClient, dogpark *pb.DogParkRequest) {
	resp, err := client.CreateDogPark(context.Background(), dogpark)
	if err != nil {
		log.Fatalf("Could not create DogPark: %v", err)
	}
	if resp.Success {
		log.Printf("A new DogPark has been added with id: %d", resp.Id)
	}
}

// getDogParks calls the RPC method GetDogParks of DogParkServer
func getDogParks(client pb.DogParkClient, filter *pb.DogParkFilter) {
	// calling the streaming API
	stream, err := client.GetDogParks(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get dogparks: %v", err)
	}
	for {
		// Receiving the stream of data 
		dogpark, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetDogParks(_) = _, %v", client, err)
		}
		log.Printf("DogPark: %v", dogpark)
	}
}
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new DogParkClient
	client := pb.NewDogParkClient(conn)

	dogpark := &pb.DogParkRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Website: "shiju@xyz.com",
		Addresses: []*pb.DogParkRequest_Address{
			&pb.DogParkRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.DogParkRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new dogpark
	createDogPark(client, dogpark)

	dogpark = &pb.DogParkRequest{
		Id:    102,
		Name:  "Irene Rose",
		Website: "irene@xyz.com",
		Addresses: []*pb.DogParkRequest_Address{
			&pb.DogParkRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new dogpark
	createDogPark(client, dogpark)
	// Filter with an empty Keyword
	filter := &pb.DogParkFilter{Keyword: ""}
	getDogParks(client, filter)
}
