package main

import (
	"fmt"

	"github.com/achimonchi/belajar_crud_postgree/config"
	"github.com/achimonchi/belajar_crud_postgree/src/modules/profile/model"
	"github.com/achimonchi/belajar_crud_postgree/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Postgree Tutorial")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	wury := model.NewProfile()
	wury.ID = "P2"
	wury.FirstName = "Jojo"
	wury.LastName = "Sinta"
	wury.Email = "jojo@gmail.com"
	wury.Password = "123456"

	// fmt.Println(wury)

	profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)
	// untuk save data
	// err = saveProfile(wury, profileRepositoryPostgres)

	// untuk update data
	// err = updateProfile(wury, profileRepositoryPostgres)

	// untuk delete data
	// err = deleteProfile(wury, profileRepositoryPostgres)

	// untuk get data by id
	// profile, err := getByID("P2", profileRepositoryPostgres)
	// fmt.Println("========================================")
	// fmt.Println(profile)

	// untuk get all data
	profiles, err := getAll(profileRepositoryPostgres)

	for _, val := range profiles {
		fmt.Println(val)
	}

}

// ini untuk save data
func saveProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Save(p)
	if err != nil {
		return err
	}

	fmt.Println("Save Success !")

	return nil
}

// ini untuk update data
func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Update(p.ID, p)
	if err != nil {
		return err
	}

	fmt.Println("Update Success !")

	return nil
}

func deleteProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Delete(p.ID)
	if err != nil {
		return err
	}

	fmt.Println("Delete Success !")

	return nil
}

// ini untuk get data by id
func getByID(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profle, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return profle, nil
}

// ini untuk get all data
func getAll(repo repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
