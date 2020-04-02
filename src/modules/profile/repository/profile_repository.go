package repository

import (
	"database/sql"

	"github.com/achimonchi/belajar_crud_postgree/src/modules/profile/model"
)

type profileRepositoryPostgres struct {
	db *sql.DB
}

// membuat constructor
func NewProfileRepositoryPostgres(db *sql.DB) *profileRepositoryPostgres {
	return &profileRepositoryPostgres{db}
}

func (r *profileRepositoryPostgres) Save(profile *model.Profile) error {
	query := `
		INSERT INTO "profile" 
		("id","first_name","last_name","email","password","created_at","updated_at")
		VALUES($1, $2, $3, $4, $5, $6, $7)
		`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	// defer berarti akan mengeksekusi paling akhir
	defer statement.Close()

	_, err = statement.Exec(profile.ID, profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.CreatedAt, profile.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *profileRepositoryPostgres) Update(id string, profile *model.Profile) error {
	query := `
		UPDATE "profile" 
		SET "first_name"=$1,"last_name"=$2, "email"=$3, "password"=$4,
		"updated_at"=$5 
		WHERE "id"=$6		
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil

}
