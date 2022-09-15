package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var ConnectDB *pgx.Conn

func Connect() string {

	godotenv.Load(".env")

	dbUrl := `host=` + host() + " " + `port=` + port() + " " + `user=` + user() + " " + `password=` + password() + " " + `dbname=` + dbname() + " " + `sslmode=` + sslmode()
	var err error
	log.Println(dbUrl)
	ConnectDB, err = pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Println(err)
		return "Error Connecting to database"
	}

	return "1"

}
