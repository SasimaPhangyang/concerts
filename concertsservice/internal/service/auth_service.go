package service

import (
	"errors"
	"strings"
	"time"

	"concerts/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour
)

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (string, string, error)
	Logout(userID int64) error
	ValidateToken(token string) (int64, error)
	SetAutoWithdraw(partnerID int, enabled bool) error
}

type authService struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthService(repo repository.AuthRepository, jwtSecret string) AuthService {
	return &authService{repo: repo, jwtSecret: jwtSecret}
}

func (s *authService) Register(name, email, password string) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return errors.New("name, email, and password are required")
	}

	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.Create(name, email, string(hashedPassword))
}

func (s *authService) Login(email, password string) (string, string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := s.generateJWT(int64(user.ID), accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateJWT(int64(user.ID), refreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	if err := s.repo.StoreToken(int64(user.ID), refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) Logout(userID int64) error {
	return s.repo.DeleteToken(userID)
}

func (s *authService) ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token")
	}

	return int64(userIDFloat), nil
}

func (s *authService) generateJWT(userID int64, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

// SetAutoWithdraw will update the auto withdraw setting for the partner.
func (s *authService) SetAutoWithdraw(partnerID int, enabled bool) error {
	// เชื่อมต่อไปยัง repository เพื่อแก้ไขค่าการตั้งค่า auto_withdraw
	return s.repo.SetAutoWithdraw(partnerID, enabled)
}
