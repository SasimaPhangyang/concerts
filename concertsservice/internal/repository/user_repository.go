package repository

import (
	"database/sql"
	"errors"

	"concerts/internal/models"

	_ "github.com/lib/pq"
)

// UserRepository interface ตามหลัก DIP
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(name, email string) (*models.User, error)
	Update(id int, name, email string) (*models.User, error)
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, created_at, updated_at FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Create(name, email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at",
		name, email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(id int, name, email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(
		"UPDATE users SET name=$1, email=$2, updated_at=now() WHERE id=$3 RETURNING id, name, email, created_at, updated_at",
		name, email, id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
