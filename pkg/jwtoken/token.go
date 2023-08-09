package jwtoken

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type TokenDetail struct {
	AccessToken  string
	ExpiredToken int64
}

type AccessDetail struct {
	UserID     string
	Authorized bool
}

type claims interface {
	Valid() error
}

func CreateToken(userid string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.ExpiredToken = time.Now().Add(time.Minute * 15).Unix()
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = td.ExpiredToken

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString("Jwt.Secret")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signature method")
		}
		return []byte(viper.GetString("Jwt.Secret")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(claims)

	if !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetail, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authorized, ok := claims["authorized"].(bool)
		if !ok {
			return nil, err
		}

		userId := strconv.FormatFloat(claims["user_id"].(float64), 'f', -1, 64)

		return &AccessDetail{
			Authorized: authorized,
			UserID:     userId,
		}, nil
	}

	return nil, err
}
