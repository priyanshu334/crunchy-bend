package anime

type Service interface {
	Create(title, description string) (*Anime, error)
	List() ([]Anime, error)
	Get(id uint) (*Anime, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(title, description string) (*Anime, error) {
	a := &Anime{
		Title:       title,
		Description: description,
	}
	return a, s.repo.Create(a)
}

func (s *service) List() ([]Anime, error) {
	return s.repo.FindAll()
}

func (s *service) Get(id uint) (*Anime, error) {
	return s.repo.FindByID(id)
}
