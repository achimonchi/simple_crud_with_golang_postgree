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

// melakukan implementasi kepada function save pada Repository
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

// melakukan implementasi kepada function Update pada Repository
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

// melakukan implementasi kepada function Delete pada Repository
func (r *profileRepositoryPostgres) Delete(id string) error {
	query := `
		DELETE FROM "profile" WHERE "id"=$1
	`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// melakukan implementasi kepada function FindByID pada Repository
// karena FindByID hanya mereturn 1 buah data, maka yang akan dia return
// adalah model.Profile
func (r *profileRepositoryPostgres) FindByID(id string) (*model.Profile, error) {
	query := `
		SELECT * FROM "profile" WHERE "id"=$1
	`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var profile model.Profile

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&profile.ID, &profile.FirstName, &profile.LastName,
		&profile.Email, &profile.Password, &profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// melakukan implementasi kepada function FindAll pada Repository
// perlu diperhatikan apa yang di return.
// karena method FindAll() akan mereturn seluruh data, maka
// return nya adalah model.Profiles
func (r *profileRepositoryPostgres) FindAll() (model.Profiles, error) {
	query := `
		SELECT * FROM "profile"
	`

	var profiles model.Profiles
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var profile model.Profile

		err = rows.Scan(
			&profile.ID, &profile.FirstName, &profile.LastName,
			&profile.Email, &profile.Password, &profile.CreatedAt,
			&profile.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil

}
