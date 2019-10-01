package main

import (
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
)

func issueToken(privatePath string) (string, error) {

	privateBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = jwt.MapClaims {
		"user": "guest",
	}
	
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
	
}

/*
func validateToken(tokenString string, publicKey string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(publicKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"])
	} else {
		fmt.Println(err)
	}

}
*/
