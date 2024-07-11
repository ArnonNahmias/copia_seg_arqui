package dao

import (
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Course struct {
	ID          uint      `gorm:"primaryKey"`
	Nombre      string    `json:"nombre"`
	Dificultad  string    `json:"dificultad"`
	Precio      float64   `json:"precio"`
	Direccion   string    `json:"direccion"`
	ImageURL    string    `json:"imageURL"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	dsn := "your-username:your-password@tcp(your-database-host:your-database-port)/your-database-name?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Course{})
}
