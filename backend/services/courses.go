package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"strconv"
)

func GetCourses() ([]domain.Course, error) {
	courses, err := clients.GetCourses1()
	if err != nil {
		return nil, err
	}
	results := make([]domain.Course, 0)
	for _, course := range courses {
		results = append(results, domain.Course{
			IdCurso:    int64(course.IdCurso), // Convert int to int64
			Nombre:     course.Nombre,
			Dificultad: course.Dificultad,
			Precio:     strconv.Itoa(course.Precio), // Convert int to string
			ImageURL:   course.ImageURL,
			CreatedAt:  course.CreatedAt,
			UpdatedAt:  course.UpdatedAt,
		})
	}
	return results, nil
}

/*func Search(query string) ([]domain.Course, error) {
	trimmed := strings.TrimSpace(query)

	courses, err := clients.SelectCoursesWithFilterName(trimmed)
	if err != nil {
		return nil, fmt.Errorf("error getting courses from DB: %w", err)
	}

	results := make([]domain.Course, 0)
	for _, course := range courses {
		results = append(results, domain.Course{
			IdCurso:    int64(course.IdCurso), // Convert int to int64
			Nombre:     course.Nombre,
			Dificultad: course.Dificultad,
			Precio:     strconv.Itoa(course.Precio), // Convert int to string
			ImageURL:   course.ImageURL,
			CreatedAt:  course.CreatedAt,
			UpdatedAt:  course.UpdatedAt,
		})
	}
	return results, nil
}
	*/

func CreateCourse(course dao.Course) (dao.Course, error) {
	result := clients.DB.Create(&course)
	return course, result.Error
}

func DeleteCourse(id int) error {
	result := clients.DB.Delete(&dao.Course{}, id)
	return result.Error
}
