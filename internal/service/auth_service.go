package service

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/No-name16/InnoTaxi-User/internal/entity"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (service *Service) CreateUser(user entity.User) (int, error) {
	dt := time.Now()
	var dtResult = dt.Format("2006-01-06 15:04:05")
	user.Password = service.GeneratePasswordHash(user.Password)
	user.UpdatedAt = dtResult
	user.CreatedAt = dtResult
	return service.repo.CreateUser(user)
}
func (service *Service) GeneratePasswordHash(password string) string {
	salt := make([]byte, 8)
	rand.Read(salt)
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (service *Service) GenerateToken(phoneNumber, password string) (string, error) {
	user, err := service.repo.GetUser(phoneNumber, service.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (service *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claim, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claim are not of type *tokenClaim")
	}
	return claim.UserId, nil
}
