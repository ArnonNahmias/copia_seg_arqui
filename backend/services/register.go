package services

import (
	"backend/clients"
	"backend/dao"
)

func Register(username, password, userType string) error {
	// Hashear la contraseña
	hashedPassword := hashPassword(password)

	user := dao.Usuario{
		NombreUsuario: username,
		Contrasena:    hashedPassword,
		Tipo:          userType,
	}

	if err := clients.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}