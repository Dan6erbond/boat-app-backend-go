package security

import (
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

type JWTUtil interface {
	SignAccessToken(claims UserClaims) ([]byte, error)
}

func NewJWTUtil(secret string) *jwtUtil {
	return &jwtUtil{
		secret:            secret,
		accessTokenSigner: jwt.NewSigner(jwt.HS256, secret, 10*time.Minute),
		verifier:          jwt.NewVerifier(jwt.HS256, secret),
	}
}

var _ JWTUtil = &jwtUtil{}

type jwtUtil struct {
	secret            string
	accessTokenSigner *jwt.Signer
	verifier          *jwt.Verifier
}

func (j *jwtUtil) SignAccessToken(claims UserClaims) ([]byte, error) {
	return j.accessTokenSigner.Sign(claims)
}
