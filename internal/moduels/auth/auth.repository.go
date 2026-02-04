package auth

import (
	"github.com/priyanshu334/go-bend/internal/moduels/user"
)

type Repository interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}
type repository struct {
	userRepo user.Repository
}

func NewRepository(userRepo user.Repository) Repository {

	return &repository{userRepo: userRepo}

}

func (r *repository) Create(user *user.User) error {
	return r.userRepo.Create(user)
}

func (r *repository) FindByEmail(email string) (*user.User, error) {

	return r.userRepo.FindByEmail(email)
}
