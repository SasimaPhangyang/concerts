package repository

import (
	"database/sql"
	"errors"

	"concerts/internal/models"

	_ "github.com/lib/pq"
)

type AuthRepository interface {
	GetByEmail(email string) (*models.PartnerUser, error)
	Create(name, email, password string) error
	StoreToken(userID int64, refreshToken string) error
	DeleteToken(userID int64) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

// GetByEmail ดึงข้อมูล partner_user จาก email
func (r *authRepository) GetByEmail(email string) (*models.PartnerUser, error) {
	var u models.PartnerUser
	err := r.db.QueryRow(`
		SELECT id, name, email, password, created_at, updated_at
		FROM partner_users
		WHERE email = $1
	`, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &u, nil
}

// Create สร้าง partner_user ใหม่
func (r *authRepository) Create(name, email, password string) error {
	_, err := r.db.Exec(`
		INSERT INTO partner_users (name, email, password)
		VALUES ($1, $2, $3)
	`, name, email, password)
	return err
}

// StoreToken เก็บ refresh token ของ user
func (r *authRepository) StoreToken(userID int64, refreshToken string) error {
	_, err := r.db.Exec(`
		INSERT INTO partner_tokens (user_id, token)
		VALUES ($1, $2)
	`, userID, refreshToken)
	return err
}

// DeleteToken ลบ refresh token ของ user
func (r *authRepository) DeleteToken(userID int64) error {
	_, err := r.db.Exec(`
		DELETE FROM partner_tokens
		WHERE user_id = $1
	`, userID)
	return err
}
