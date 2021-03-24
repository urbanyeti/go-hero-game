package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urbanyeti/go-hero-game/game"
	pb "github.com/urbanyeti/go-hero-game/server/grpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "x509/localhost.crt", "The TLS cert file")
	keyFile  = flag.String("key_file", "x509/localhost.key", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
)

type gameWorldServer struct {
	pb.UnimplementedGameWorldServer
	g *game.Game
}

func (s *gameWorldServer) Initialize() {
	s.g = &game.Game{}
	s.g.Initialize()
}

func (s *gameWorldServer) GetRandomItem(context.Context, *pb.ItemRequest) (*pb.Item, error) {
	item := s.g.Items.GetRandomItem()

	return pb.PackItem(item), nil
}

func newServer() *gameWorldServer {
	s := &gameWorldServer{}
	s.Initialize()
	return s
}

func main() {
	os.Chdir("..")
	game := game.Game{}
	game.Initialize()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGameWorldServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
