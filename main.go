package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urbanyeti/go-hero-game/game"
)

const maxTurns int = 100

func setLogging() *os.File {
	var filename string = "logfile.log"
	// Create the log file if doesn't exist
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)

	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}

	return f

}

func main() {
	f := setLogging()
	defer f.Close()
	game := game.Game{}
	game.Initialize()

	for i := 0; i < maxTurns; i++ {
		gameOver := game.PlayTurn()
		if gameOver {
			break
		}
	}
}
