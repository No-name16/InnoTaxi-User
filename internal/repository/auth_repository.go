package repository

import (
	"github.com/No-name16/InnoTaxi-User/internal/entity"
)

func (repo *Repository) CreateUser(user entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, phonenumber, email, password, createdat, updatedat) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := repo.db.QueryRow(query, user.Name, user.PhoneNumber, user.Email, user.Password, user.CreatedAt,
		user.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *Repository) GetUser(phonenumber, password string) (entity.User, error) {
	var user entity.User
	query := "SELECT id FROM users  WHERE phonenumber=$1 AND password=$2"
	err := repo.db.Get(&user, query, phonenumber, password)
	return user, err
}

func (repo *Repository) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	query := "SELECT * FROM users WHERE id=$1"
	err := repo.db.Get(&user, query, id)
	return user, err
}
