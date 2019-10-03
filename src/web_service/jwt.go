package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

func issueToken(privatePath string) (string, error) {

	// ReadFileは、指定のファイルを読み込んで、byte配列で返却する。
	privateBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return "", err
	}

	// jwt.ParseRSAPrivateKeyFromPEMは、PEMエンコードされた秘密鍵（PKCS1,PKCS8）をパースして、byte配列で返却する。
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return "", err
	}

	// jwt.Newは、署名アルゴリズムを指定してトークンを作成する。
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// JWTのペイロードを設定する。
	token.Claims = jwt.MapClaims {
		// 予約されていないデータ
		"user": "guest",
		// 以降は、予約されたデータ（トークン発行日時、有効日時）
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(1 * time.Minute).Unix(),
	}
	
	// SignedStringは、秘密鍵を元にトークン文字列を作成する。
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
	
}

func validateToken(tokenString string, publicPath string) (bool, error) {

	publicBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return false, err
	}

	// jwt.ParseRSAPublicKeyFromPEMは、PEMエンコードされた公開鍵を読み込んで、byte配列で返却する。
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return false, err
	}

	// jwt.Parseは、トークン文字列をパース＆検証して、トークンを返却する。
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムがRS256か確認する。
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	// トークンからペイロードが取得できてかつ、トークンが有効か確認する。
	// 型アサーションによる型変換を利用。
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"])
		return true, nil
	}

	return false, err

}
