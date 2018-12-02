package jwt

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Claims contains the fields encoded in the JWT token
type Claims struct {
	ID        int    `json:"sub"`
	Username  string `json:"name"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

// Encode encodes a JWT token containing the passed in claims and
// using the passed in secret
func Encode(claims *Claims, secret string) (string, error) {
	now := time.Now()
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(10 * time.Minute).Unix()

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	mapClaims := jwt.MapClaims{}
	err = json.Unmarshal(jsonClaims, &mapClaims)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(secret))
}

// GetClaims returns the JWT claims from the Echo context
func GetClaims(c echo.Context) (Claims, error) {
	claims := Claims{}

	token := c.Get("user").(*jwt.Token)
	mapClaims := token.Claims.(jwt.MapClaims)

	jsonClaims, err := json.Marshal(mapClaims)
	if err != nil {
		return claims, err
	}
	err = json.Unmarshal(jsonClaims, &claims)
	if err != nil {
		return claims, err
	}
	return claims, nil
}
