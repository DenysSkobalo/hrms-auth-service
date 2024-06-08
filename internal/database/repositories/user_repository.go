package repositories

import (
	"database/sql"
	"fmt"
	"hrms-auth-service/internal/database/utils"
	"hrms-auth-service/internal/models"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)

	IsEmailExists(email string) bool
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, password_hash, first_name, last_name, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.DB.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, email, password_hash, first_name, last_name, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) IsEmailExists(email string) (bool, error) {
	return utils.IsEmailExists(r.DB, email)
}
