package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	GenerateToken(id string, username string, role string ) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	Test()
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "my_secret_key",
		tokenTtl: 1000,
	}
}

type jwtService struct {
	secretKey string
	tokenTtl  int64
}

type TokenClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (jwtSry *jwtService) GenerateToken(id string, username string, role string ) (string, error) {

	claims := TokenClaims{
		Id:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute  * time.Duration(jwtSry.tokenTtl)).Unix(),
			Issuer:    id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSry.secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, err
}

func (jwtSry *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSry.secretKey), nil
	})

	return token, err
}


func (jwtSry *jwtService) Test() {

	tokenString, err := jwtSry.GenerateToken("1", "22", "ew")
	fmt.Println("err ", err)
	fmt.Println("tokenString ", tokenString)

	token, err := jwtSry.ValidateToken(tokenString)
	fmt.Println("err ", err)
	fmt.Println("token ", token)

	fmt.Println("token.Valid ", token.Valid)
}

