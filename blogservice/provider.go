package blogservice

import (
	"github.com/google/wire"
	"github.com/koetsuserizawa/wire_example/userservice"
)

var BlogServiceSet = wire.NewSet(
	NewBlogService,
	NewBlogRepository,
	NewDb,
	userservice.NewUserService,
	userservice.NewUserRepository,
	userservice.NewDb,
)
