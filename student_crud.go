// For working with Database

package main

import "gorm.io/gorm"

// * C
func createStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

// * R
func getAllStudent(db *gorm.DB) ([]Student, error) {
	// create variable to stored fetching data
	var student []Student
	result := db.Find(&student)
	return student, result.Error
}

// R if data to find is PK
func getStudentByID(db *gorm.DB, id int) (*Student, error) {
	// create variable to stored fetching data
	var student Student
	result := db.First(&student, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

// R if data to find is not PK
func getStudentByName(db *gorm.DB, StuName string) ([]Student, error) {
	// create variable to stored fetching data
	var students []Student

	result := db.Where("stu_name LIKE ?", "%"+StuName+"%").Find(&students)

	return students, result.Error
}

func getStudentByLevel(db *gorm.DB, StuLevel int) ([]Student, error) {
	// create variable to stored fetching data
	var students []Student

	result := db.Where("stu_level = ?", StuLevel).Find(&students)

	return students, result.Error
}

// * U
func updateStudent(db *gorm.DB, id int, student *Student) error {
	return db.Model(&Student{}).Where("id = ?", id).Updates(student).Error
}

// * D
func deleteStudent(db *gorm.DB, id int) error {
	return db.Delete(&Student{}, id).Error
}
