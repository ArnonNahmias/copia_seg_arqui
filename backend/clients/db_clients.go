package clients

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"backend/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:admin@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:58005800@tcp(127.0.0.1:3306)/proyecto1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")

	Migrate()
	SeedDB()
}

func Migrate() {
	log.Println("Migrating database...")
	DB.AutoMigrate(&dao.Usuario{}, &dao.Course{}, &dao.Subscription{})
}

func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SeedDB() {
	log.Println("Seeding database...")

	// Hashear las contraseñas
	adminPassword := hashPassword("admin")
	userPassword := hashPassword("user")

	admin := dao.Usuario{NombreUsuario: "admin", Contrasena: adminPassword, Tipo: "admin"}
	user := dao.Usuario{NombreUsuario: "user", Contrasena: userPassword, Tipo: "normal"}

	DB.FirstOrCreate(&admin, dao.Usuario{NombreUsuario: "admin"})
	DB.FirstOrCreate(&user, dao.Usuario{NombreUsuario: "user"})

	log.Println("Seeding database...")
	cursos := []dao.Course{
		{Nombre: "Ingles B2", Dificultad: "Medio", Precio: 45, Direccion: "José Roque Funes 1511 X5000ABE Córdoba",
			Rating: 4.5, Categoria: "Idiomas", Archivo1: "https://www.youtube.com/watch?v=6PtflpPVhJQ", Archivo2: " ",
			Comentarios: "Curso interesante para aquellos que desean la mejor experiencia en nivel de ingles B2"},
	}
	for _, curso := range cursos {
		DB.Create(&curso)
	}
	log.Println("Database seeded successfully")
}

func GetCourses1() ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectCoursesWithFilterName(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Where("Nombre LIKE ? ", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}
