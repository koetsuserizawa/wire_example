package userservice

type UserServiceInterface interface {
	Register(string, string) error
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

func (s *UserService) Register(id string, name string) error {
	return s.repo.Register(id, name)
}

func (s *UserService) Get(userID string) (*User, error) {
	return s.repo.Get(userID)
}
