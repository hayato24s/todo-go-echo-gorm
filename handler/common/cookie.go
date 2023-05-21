package common

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var tokenName string = os.Getenv("TOKEN_NAME")
var tokenKey string = os.Getenv("TOKEN_KEY")
var tokenMaxAge int

func init() {
	tokenMaxAge, _ = strconv.Atoi(os.Getenv("TOKEN_MAX_AGE"))
}

func GetUserIDFromCookie(c echo.Context) (userID uuid.UUID, err error) {
	cookie, err := c.Cookie(tokenName)
	if err != nil {
		return uuid.Nil, err
	}
	tokenString := cookie.Value

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected loging method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err = uuid.Parse(claims["userID"].(string))
	}
	return userID, err
}

func SetUserIDInCookie(c echo.Context, userID uuid.UUID) error {
	// Create a new token object, specifying loging method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID.String(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     tokenName,
		Value:    tokenString,
		MaxAge:   tokenMaxAge,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	return nil
}

func ClearUserIDInCookie(c echo.Context) {
	cookie := &http.Cookie{
		Name:     tokenName,
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}
