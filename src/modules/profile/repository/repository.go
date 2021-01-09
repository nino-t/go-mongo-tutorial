package repository

import (
	"github.com/nino-t/go-mongo-tutorial/src/modules/profile/model"
)

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindByAll() (model.Profiles, error)
}
