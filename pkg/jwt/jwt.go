package jwt

import (
  "fmt"
  "time"
  "os"
  . "github.com/dgrijalva/jwt-go"
)

type customClaims struct {
  Username string `json:"username"`
  IsUser bool `json:"isUser"`
  StandardClaims
}

func GenerateJWT (username string, isUser bool) (string, error) {

  claims := customClaims{
    username,
    isUser,
    StandardClaims{
      ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
      Issuer: "capricorn",
      IssuedAt: time.Now().Unix(),
    },
  }

  secret := os.Getenv("CAPRICORN_JWT_SECRET")
  token := NewWithClaims(SigningMethodHS256, claims)
  t, err := token.SignedString([]byte(secret))

  if err != nil {
    return "", err
  }

  return t, nil

}

func ValidateJWT(encodedToken string) (*Token, error){

  secret := os.Getenv("CAPRICORN_JWT_SECRET")

  return Parse(encodedToken, func (token *Token) (interface{}, error) {
    if _, isvalid := token.Method.(*SigningMethodHMAC); !isvalid {
      return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])
    }
    return secret, nil
  })

}

