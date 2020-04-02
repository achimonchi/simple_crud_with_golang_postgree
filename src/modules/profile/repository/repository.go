package repository

import (
	"github.com/achimonchi/belajar_crud_postgree/src/modules/profile/model"
)

type ProfileRepository interface {
	// berisi method yang akan digunakan
	// formatnya adalah
	// namaMethod(parameter) return_value

	// akan menerima data bertipe Profile, dan mereturn data bertipe error
	Save(*model.Profile) error

	// akan menerima data bertipe string, dan Profile, dan mereturn data bertipe error
	Update(string, *model.Profile) error

	// akan menerima data bertipe string, dan mereturn data bertipe error
	Delete(string) error

	// akan menerima data bertipe string, dan mereturn data bertipe profile dan error
	FindByID(string) (*model.Profile, error)

	// akan mereturn data bertipe profile dan error
	FindAll() (model.Profile, error)
}
