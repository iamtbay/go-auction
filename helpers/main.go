package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaims struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

// CREATE JWT
func CreateJWT(userInfo *models.UserInfoModel) (string, error) {
	claims := JWTClaims{
		userInfo.ID,
		userInfo.Username,
		userInfo.Email,
		jwt.RegisteredClaims{
			Issuer:    userInfo.ID.Hex(),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// PARSE JWT
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

func GetUserClaimsFromJWT(c *fiber.Ctx) (*JWTClaims, error) {
	tokenString := c.Cookies("accessToken")
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
