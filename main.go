package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v", user, host, port, dbName)

	http.Handle("/articles", initializeBlogServiceHandler(dsn))
	http.Handle("/users", initializeUserServiceHandler(dsn))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
