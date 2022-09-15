package database

import "os"

func host() string {
	return os.Getenv("HOST")
}

func port() string {
	return os.Getenv("PORT")
}

func user() string {
	return os.Getenv("USER_NAME")
}

func password() string {
	return os.Getenv("PASSWORD")
}

func dbname() string {
	return os.Getenv("DBNAME")
}

func sslmode() string {
	return os.Getenv("SSLMODE")
}
