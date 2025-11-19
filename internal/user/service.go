package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repository
	jwtSecret []byte
	
}

func NewService(repo *Repository, jwtSecret []byte) *Service {
	return &Service{repo, jwtSecret}
}



func (s *Service) Register(email, password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(email, string(hashed))
}


func (s *Service) Login(email, password string) (string, error) {
	u, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// IMPORTANT: use PasswordHash (matches DB)
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(s.jwtSecret)
}
