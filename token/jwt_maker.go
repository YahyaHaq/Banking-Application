package token

// import (
// 	"fmt"
// 	"time"

// )

// const minSecretKeySize = 32

// type JWTMaker struct {
// 	secretKey string
// }

// func NewJWTMaker(secretKey string) (Maker, error) {
// 	if len(secretKey) < minSecretKeySize {
// 		return nil, fmt.Errorf("Invalid key size: must be atleast %d characters", minSecretKeySize)
// 	}

// 	return &JWTMaker{secretKey}, nil
// }

// func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
// 	payload, err := NewPayload(username, duration)
// 	if err != nil {
// 		return "", err
// 	}

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
// }
