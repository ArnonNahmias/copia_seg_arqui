package services

import (
	"errors"
	"time"

	"backend/clients"
	"backend/dao"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Login(username, password string) (string, int, string, error) {
	var user dao.Usuario
	if err := clients.DB.Where("Nombre_usuario = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", 0, "", errors.New("Invalid credentials")
		}
		return "", 0, "", err
	}

	// Hashear la contraseña ingresada y compararla con la almacenada
	hashedPassword := hashPassword(password)
	if user.Contrasena != hashedPassword {
		return "", 0, "", errors.New("Invalid credentials")
	}

	// Generar JWT
	token, err := GenerateJWT(username)
	if err != nil {
		return "", 0, "", err
	}

	return token, user.IdUsuario, user.Tipo, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}