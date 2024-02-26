package auth

import (
	"time"
	"todos/api"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func authorize(token string) (*api.User, error) {
	tok, err := jwt.ParseWithClaims(token, &userTokenClaims{}, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", tok.Header["alg"])
		}
		return []byte("-foobar-"), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token")
	}
	if claims, ok := tok.Claims.(*userTokenClaims); ok && tok.Valid {
		return lookup(claims.UserEmail), nil
	}
	return nil, errors.New("invalid token")
}

type userTokenClaims struct {
	jwt.StandardClaims
	UserEmail string
}

func token(u *api.User) string {
	claims := userTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "go-adventure",
		},
		UserEmail: u.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("-foobar-"))
	return tokenString
}
