package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"myFinanceTask/internal/handler/rest"
	"time"
)

const (
	salt       = "universe517"
	signingKey = "worldisyours758"
)

type authService struct {
	repo Authorization
}

func NewAuthService(repo Authorization) *authService {
	return &authService{repo: repo}
}

func (s *authService) CreateUser(user rest.UserDTO) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *authService) IsAdmin(id int64) bool {
	return s.repo.IsAdmin(id)
}

//Jwt tokens

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (s *authService) GenerateToken(name, password string) (string, error) {
	log.Println("Generating auth token")
	password = generateHashPassword(password)
	user, err := s.repo.GetUser(name, password)
	if err != nil {
		return "", err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.ID,
	}).SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.New("error with generate token: " + err.Error())
	}
	return token, err
}

func (s *authService) ParseToken(accessToken string) (int64, error) {
	log.Println("Parsing access token")
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
