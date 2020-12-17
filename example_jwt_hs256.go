package main

import (
	"os"
	"time"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)




func GenerateJWT() (string, error){
	// Load the .env file in the current directory
	godotenv.Load()
	// get the secret from .env
	var secretKey2 = []byte(os.Getenv("SECRET_KEY"))
	// Generate a token based on the secret key
	// https://sodocumentation.net/go/topic/10161/jwt-authorization-in-go
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": "true",
		"user":"dl_copy_cli",
		"exp": time.Now().Add(time.Minute * 30).Unix(),

	})
	tokenString, err := token.SignedString(secretKey2)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}


func main() {



	tokenString, err := GenerateJWT()
	if err != nil {
			fmt.Println("Error creating tokenString")
		}
	fmt.Println(tokenString)
}