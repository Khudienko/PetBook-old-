package models

type Pet struct {
	ID          int     `db:"user_id"`
	Name        string  `db:"name"`
	PetType     string  `db:"animal_type"`
	Breed       string  `db:"breed"`
	Age         string    `db:"age"`
	Weight      string `db: "weight"`
	Gender      string  `db:"gender"`
	Description string  `db:"description"`

}
