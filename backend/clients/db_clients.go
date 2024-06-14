package clients

import (
	"backend/dao"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:DeanF9360@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
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

func SeedDB() {
	log.Println("Seeding database...")
	admin := dao.Usuario{NombreUsuario: "admin", Contrasena: "admin", Tipo: "admin"}
	user := dao.Usuario{NombreUsuario: "user", Contrasena: "user", Tipo: "normal"}

	DB.FirstOrCreate(&admin, dao.Usuario{NombreUsuario: "	admin"})
	DB.FirstOrCreate(&user, dao.Usuario{NombreUsuario: "user"})

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
