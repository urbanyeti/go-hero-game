# go-hero-game
log "github.com/sirupsen/logrus"

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/game_world.proto