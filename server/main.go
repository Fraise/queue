package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"server/database"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := database.New("database.sqlite")

	if err != nil {
		log.Fatal().Err(err).Msg("error opening database")
	}

	_ = db
}
