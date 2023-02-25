package auth

import (
	"fmt"
	"os"
	"time"
	
	jwt "github.com/dgrijalva/jwt-go"
)

type signedDetails struct {
	ID      uint 
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")




func GenerateAccessToken(id uint,role string) (signedToken string, err error) {
	claims := &signedDetails{
		ID:      id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(3)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))


	return token,err
}




func ValidateToken(signedToken string) (claims *signedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&signedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return claims, msg
	}

	claims, ok := token.Claims.(*signedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		return claims, msg
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		return claims, msg
	}

	return claims, msg
}

