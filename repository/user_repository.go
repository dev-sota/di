package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	FindById(identifier int)
}

type userRepository struct {
	db sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{*db}
}

func (ur *userRepository) FindById(identifier int) {
	row, _ := ur.db.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
	var id int
	var firstName string
	var lastName string
	row.Next()
	_ = row.Scan(&id, &firstName, &lastName)
	fmt.Println(id, firstName, lastName)
}
