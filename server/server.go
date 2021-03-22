package main

import (
	context "context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urbanyeti/go-hero-game/game"
	server "github.com/urbanyeti/go-hero-game/server/grpc"
	grpc "google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type gameWorldServer struct {
	server.UnimplementedGameWorldServer
	g *game.Game
}

func (s *gameWorldServer) Init() {
	s.g = &game.Game{}
	s.g.Init()
}

func (s *gameWorldServer) GetRandomItem(context.Context, *server.ItemRequest) (*server.Item, error) {
	item := s.g.Items.GetRandomItem()
	v, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return nil, err
	}

	return &server.Item{Data: v}, nil
}

func newServer() *gameWorldServer {
	s := &gameWorldServer{}
	s.Init()
	return s
}

func main() {
	os.Chdir("..")
	game := game.Game{}
	game.Init()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server.RegisterGameWorldServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
