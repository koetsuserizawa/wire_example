package blogservice

type BlogServiceInterface interface {
	Read(string) ([]Article, error)
	Write(string, string) error
}

type BlogService struct {
	repo BlogRepositoryInterface
}

func NewBlogService(r BlogRepositoryInterface) BlogServiceInterface {
	return &BlogService{
		repo: r,
	}
}

func (s *BlogService) Read(userID string) ([]Article, error) {
	return s.repo.Read(userID)
}

func (s *BlogService) Write(userID string, text string) error {
	return s.repo.Write(userID, text)
}
