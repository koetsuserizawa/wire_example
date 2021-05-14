package userservice

import "github.com/google/wire"

var UserServiceSet = wire.NewSet(
	NewUserService,
	NewUserRepository,
	NewDb,
)
