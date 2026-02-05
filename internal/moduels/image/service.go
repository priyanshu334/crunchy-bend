package image

type Service interface {
	AddImage(url, ownerType string, ownerID uint) error
	GetImages(ownerType string, ownerID uint) ([]Image, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AddImage(url, ownerType string, ownerID uint) error {
	img := &Image{
		URL:       url,
		OwnerType: ownerType,
		OwnerID:   ownerID,
	}
	return s.repo.Create(img)
}

func (s *service) GetImages(OwnerType string, ownerId uint) ([]Image, error) {
	return s.repo.FindByOwner(OwnerType, ownerId)
}
