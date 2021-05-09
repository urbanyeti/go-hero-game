module github.com/urbanyeti/go-hero-game

go 1.16

replace github.com/urbanyeti/go-hero-game/math => ./math

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2
	github.com/hajimehoshi/ebiten v1.12.12
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
)
