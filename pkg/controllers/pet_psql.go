package controllers

import (
	"PetBook/pkg/models"
	//	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	//"log"
)

type PetStore struct {
	DB *sqlx.DB
}

type PetStorer interface {
	GetPet(pet *models.Pet) error
	RegisterPet(pet *models.Pet) error
}

// func (c *UserStore) GetU() ([]models.User, error) {
// 	rows, err := c.DB.Query("select * from users")
// 	logErr(err)
// 	defer rows.Close()
// 	users := []models.User{}
// 	err = sqlx.StructScan(rows, &users)
// 	if err != nil {
// 		logErr(err)
// 		return users, fmt.Errorf("cannot scan users from db: %v", err)
// 	}
// 	return users, nil
// }

func (c *PetStore) GetPet(pet *models.Pet) error {
	err := c.DB.QueryRowx("select * from pets where user_id=$1", pet.ID).StructScan(pet)
	if err != nil {
		logErr(err)
		return fmt.Errorf("cannot scan pet from db: %v", err)
	}
	return nil
}

func (c *PetStore) RegisterPet(pet *models.Pet) error {
	_, err := c.DB.Exec("insert into pets (user_id, name, animal_type,breed,age,weight, gender) values ($1, $2, $3, $4, $5, $6, $7)",
		pet.ID, pet.Name, pet.PetType, pet.Breed, pet.Age, pet.Weight, pet.Gender)
	if err != nil {
		return fmt.Errorf("cannot affect rows in pets in db: %v", err)
	}
	return nil
}
