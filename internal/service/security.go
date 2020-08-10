package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/seinyan/go-rest-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SecurityService interface {
	GeneratePasswordHash(password string) (string, error)
	compareHashAndPassword(password string, passwordHash string) error
	IsAuthenticated(user models.User) bool
	GenerateToken(id string, username string, role string ) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	TestJWT()
}

type securityService struct {
	secretKey string
	tokenTtl  int64
}

type TokenClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func NewSecurityService() SecurityService {
	return &securityService{
		secretKey: "my_secret_key",
		tokenTtl: 1000,
	}
}

func (s securityService) GeneratePasswordHash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func (s securityService) IsAuthenticated(user models.User) bool {
	if err := s.compareHashAndPassword(user.Password, user.PasswordHash); err != nil {
		return false
	}
	return true
}

func (s securityService) compareHashAndPassword(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}


// *** JWT *** //

func (s *securityService) GenerateToken(id string, username string, role string ) (string, error) {

	claims := TokenClaims{
		Id:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute  * time.Duration(s.tokenTtl)).Unix(),
			Issuer:    id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, err
}

func (s *securityService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	return token, err
}

// Test JWT token
func (s *securityService) TestJWT() {

	tokenString, err := s.GenerateToken("1", "22", "ew")
	fmt.Println("err ", err)
	fmt.Println("tokenString ", tokenString)

	token, err := s.ValidateToken(tokenString)
	fmt.Println("err ", err)
	fmt.Println("token ", token)

	fmt.Println("token.Valid ", token.Valid)
}

