package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
	"os"
)

const schema = `
CREATE TABLE topic (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT
);

CREATE UNIQUE INDEX topic_id_unique_index ON topic (id);

CREATE TABLE message (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp DATETIME,
	content TEXT,
	uuid TEXT,
	topic_id INTEGER,
	CONSTRAINT fk_topic
		FOREIGN KEY(topic_id)
	    REFERENCES topics(id),
	CONSTRAINT uuid_unique UNIQUE (uuid)
);

CREATE UNIQUE INDEX message_id_unique_index ON message (id);
`

type Database struct {
	reader *sqlx.DB
	writer *sqlx.DB
}

func New(path string) (*Database, error) {
	db := new(Database)
	var err error

	isNewDb := false

	if _, err = os.Stat(path); os.IsNotExist(err) {
		isNewDb = true
	}

	log.Debug().Msg("creating database writer")

	db.writer, err = sqlx.Connect("sqlite", "file:"+path)
	if err != nil {
		return nil, fmt.Errorf("error creating database writer: %w", err)
	}

	db.writer.SetMaxOpenConns(1)

	if isNewDb {
		log.Debug().Msg("new database, initializing settings and tables")

		_, err = db.writer.Exec("PRAGMA journal_mode = 'wal';")
		if err != nil {
			return nil, fmt.Errorf("error enabling WAL: %w", err)
		}

		_, err = db.writer.Exec(schema)
		if err != nil {
			return nil, fmt.Errorf("error creating database schema: %w", err)
		}
	}

	log.Debug().Msg("creating database reader")

	db.reader, err = sqlx.Connect("sqlite", "file:"+path+"?mode=ro")
	if err != nil {
		return nil, fmt.Errorf("error creating database reader: %w", err)
	}

	return db, nil
}
