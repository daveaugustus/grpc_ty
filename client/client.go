package main

import (
	"context"
	"fmt"
	pb "grpc_ty/gen/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("cannot dial to localhost")
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello World!"})
	if err != nil {
		log.Printf("no response from localhost")
	}

	fmt.Println("Yayy server sent us: ", resp.Msg)
}
