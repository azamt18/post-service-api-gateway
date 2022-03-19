package main

import (
	"context"
	"fmt"
	apiClient "github.com/azamt18/post-service-grpc-api-gateway/api-gateway/api/goclient/v1"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

const apiSvc = "localhost:9090"

func main() {
	conn, err := grpc.DialContext(context.TODO(), apiSvc, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	api := apiClient.NewPostServiceClient(conn)

	res, err := api.ReadPost(context.Background(), &apiClient.ReadPostRequest{PostId: "123"})
	if err != nil {
		panic(err)
	}

	resp, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		panic(err)
	}

	fmt.Printf(resp)
}
