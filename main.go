package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/koetsuserizawa/wire_example/blogservice"
	"github.com/koetsuserizawa/wire_example/handlers"
	"github.com/koetsuserizawa/wire_example/userservice"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v", user, host, port, dbName)
	userCnf := userservice.Config{DSN: dsn}
	bCnf := blogservice.Config{DSN: dsn}

	db, err := userservice.NewDb(userCnf)
	d, err := blogservice.NewDb(bCnf)
	if err != nil {
		panic(err)
	}
	ur := userservice.NewUserRepository(db)
	us := userservice.NewUserService(ur)
	uh := handlers.NewUserServiceHandler(us)

	br := blogservice.NewBlogRepository(d)
	bs := blogservice.NewBlogService(br)
	bh := handlers.NewBlogServiceHandler(bs, us)

	http.Handle("/articles", bh)
	http.Handle("/users", uh)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
