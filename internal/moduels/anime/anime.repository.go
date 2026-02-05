package anime

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(anime *Anime) error
	FindAll() ([]Anime, error)
	FindByID(id uint) (*Anime, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(anime *Anime) error {
	return r.db.Create(anime).Error
}

func (r *repository) FindAll() ([]Anime, error) {
	var list []Anime
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) FindByID(id uint) (*Anime, error) {
	var anime *Anime
	err := r.db.Find(&anime, id).Error
	return anime, err
}
