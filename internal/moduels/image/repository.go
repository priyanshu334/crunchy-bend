package image

import "gorm.io/gorm"

type Repository interface {
	Create(img *Image) error
	FindByOwner(ownerType string, ownerId uint) ([]Image, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
func (r *repository) Create(img *Image) error {
	return r.db.Create(img).Error
}

func (r *repository) FindByOwner(ownerType string, ownerId uint) ([]Image, error) {
	var images []Image
	err := r.db.Where("ownerType=? AND ownerId=?", ownerType, ownerId).Find(&images).Error
	return images, err
}
