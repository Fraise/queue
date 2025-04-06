package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"server/database"
)

//go:generate protoc --proto_path=../proto --go_out=server/grpc --go-grpc_out=server/grpc queue.proto
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := database.New("database.sqlite")

	if err != nil {
		log.Fatal().Err(err).Msg("error opening database")
	}

	_ = db
}
