package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urbanyeti/go-hero-game/character"
	"github.com/urbanyeti/go-hero-game/game"
	server "github.com/urbanyeti/go-hero-game/server/grpc"
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
	server.UnimplementedGameWorldServer
	g *game.Game
}

func (s *gameWorldServer) Initialize() {
	s.g = &game.Game{}
	s.g.Initialize()
}

func (s *gameWorldServer) GetRandomItem(context.Context, *server.ItemRequest) (*server.Item, error) {
	item := s.g.Items.GetRandomItem()

	return itemResponse(item), nil
}

func itemResponse(item *character.Item) *server.Item {
	r := &server.Item{}
	r.ID = item.ID()
	r.Name = item.Name()
	r.Desc = item.Desc()
	r.Tags = item.Tags
	r.Stats = make(map[string]int32, len(item.Stats))
	for k, v := range item.Stats {
		r.Stats[k] = int32(v)
	}

	return r
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
	server.RegisterGameWorldServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
