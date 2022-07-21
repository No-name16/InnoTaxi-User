package repository

import (
	"fmt"
	"github.com/No-name16/InnoTaxi-User/internal/entity"
)

func (repo *Repository) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, phonenumber, email, password, createdat, updatedat) values " +
		"($1, $2, $3, $4, $5, $6) RETURNING id ")
	row := repo.db.QueryRow(query, user.Name, user.PhoneNumber, user.Email, user.Password, user.CreatedAt,
		user.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
