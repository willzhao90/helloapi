package main

import (
	"context"
	"fmt"

	//"log"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	pb "github.com/willzhao90/hellobackend/out"
	"google.golang.org/grpc"
)

const (
	rpcendpoint = "127.0.0.1:8030"
)

type ApiServer struct {
	rpcClient pb.HelloServiceClient
}

func main() {
	conn, err := grpc.Dial(rpcendpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	server := &ApiServer{
		rpcClient: c,
	}

	http.HandleFunc("/", server.HelloServer)
	http.ListenAndServe(":8080", nil)
}

func (s *ApiServer) HelloServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := s.rpcClient.GetHello(ctx, &pb.GetHelloRequest{
		Name: r.URL.Path[1:],
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", res.Name)
	//io.WriteString(w, res.Name)
}

// func getRpcClient() *pb.HelloServiceClient {
// 	conn, err := grpc.Dial(rpcendpoint, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewHelloServiceClient(conn)
// 	return &c
// }
