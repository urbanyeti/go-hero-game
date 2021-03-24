package main

import (
	"context"
	"flag"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	pb "github.com/urbanyeti/go-hero-game/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls        = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("ca_file", "../x509/localhost.crt", "The file containing the cert file")
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func getRandomItem(client pb.GameWorldClient, req *pb.ItemRequest) {
	log.WithFields(log.Fields{"request": req, "client": client}).Info("getting random item")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	item, err := client.GetRandomItem(ctx, req)
	if err != nil {
		log.WithFields(log.Fields{"client": client, "error": err}).Fatal("GetRandomItem")
	}
	i := pb.UnpackItem(item)
	log.Println(i)
}

func getMonsters(client pb.GameWorldClient, req *pb.MonsterRequest) {
	log.WithFields(log.Fields{"request": req, "client": client}).Info("getting monsters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetMonsters(ctx, req)
	if err != nil {
		log.WithFields(log.Fields{"client": client, "error": err}).Fatal("GetMonsters")
	}
	for {
		monster, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.WithFields(log.Fields{"client": client, "error": err}).Fatal("GetMonsters")
		}
		log.WithFields(log.Fields{"monster": monster}).Info("GetMonsters")
	}

}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*certFile, "")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGameWorldClient(conn)
	getRandomItem(client, &pb.ItemRequest{LoopNumber: 1, Level: 1})
	getMonsters(client, &pb.MonsterRequest{LoopNumber: 1, Level: 1})
}
