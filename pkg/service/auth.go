package service

import (
	todo "Pet-project-ToDoApp"
	"Pet-project-ToDoApp/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "csiugfioieu7syf8r9w8er9we8fy"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
