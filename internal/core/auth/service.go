package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"myFinanceTask/internal/handler/rest"
	"time"
)

const (
	salt       = "universe517"
	signingKey = "worldisyours758"
)

type AuthService struct {
	repo Authorization
}

func NewAuthService(repo Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest.UserDTO) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

//Jwt tokens

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (s *AuthService) GenerateToken(name, password string) (string, error) {
	password = s.generateHashPassword(password)
	user, err := s.repo.GetUser(name, password)
	if err != nil {
		return "", err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.ID,
	}).SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.New("error with generate token: " + err.Error())
	}
	return token, err
}

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) { // Ссылка !!!!! tokenClaims
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid { //Ссылка tokenClaims!!!!!
		return claims.UserId, nil
	} else {
		return 0, errors.New("invalid token")
	}
}
