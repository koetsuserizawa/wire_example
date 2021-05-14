package userservice

type UserServiceInterface interface {
	Register(string) (*User, error)
	Get(string) (*User, error)
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(r UserRepositoryInterface) UserRepositoryInterface {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Register(name string) (*User, error) {
	return s.repo.Register(name)
}

func (s *UserService) Get(userID string) (*User, error) {
	return s.repo.Get(userID)
}
