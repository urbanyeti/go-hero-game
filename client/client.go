package main

import (
	"context"
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
	pb "github.com/urbanyeti/go-hero-game/server/grpc"
	"google.golang.org/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func getRandomItem(client pb.GameWorldClient, req *pb.ItemRequest) {
	log.WithFields(log.Fields{"request": req, "client": client}).Info("getting random item")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	item, err := client.GetRandomItem(ctx, req)
	if err != nil {
		log.WithFields(log.Fields{"client": client, "error": err}).Fatal("GetRandomItem")
	}
	log.Println(item)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGameWorldClient(conn)
	getRandomItem(client, &pb.ItemRequest{LoopNumber: 1, Level: 1})
}
