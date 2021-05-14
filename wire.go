// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/koetsuserizawa/wire_example/blogservice"
	"github.com/koetsuserizawa/wire_example/handlers"
	"github.com/koetsuserizawa/wire_example/userservice"
)

func initializeUserServiceHandler(dsn string) *handlers.UserServiceHandler {
	wire.Build(userservice.UserServiceSet, handlers.NewUserServiceHandler)
	return nil
}

func initializeBlogServiceHandler(dsn string) *handlers.BlogServiceHandler {
	wire.Build(blogservice.BlogServiceSet, handlers.NewBlogServiceHandler)
	return nil
}
