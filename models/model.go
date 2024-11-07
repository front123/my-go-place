package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 学生表结构
type Student struct {
	gorm.Model // 内嵌gorm.Model，包含ID，CreatedAt，UpdatedAt，DeletedAt等字段
	Name       string
	Age        uint
	Class      string
}

// 课程表结构
type Course struct {
	gorm.Model
	Name    string
	Teacher string
	Room    string
}

// 学生课程成绩表结构
type Score struct {
	gorm.Model
	StudentID uint // 外键，关联学生表的ID
	CourseID  uint // 外键，关联课程表的ID
	Score     int  // 分数
	// Student   Student `gorm:"foreignKey:StudentID"` // 学生信息
	// Course    Course  `gorm:"foreignKey:CourseID"`  // 课程信息
}

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "root:pwd123@tcp(127.0.0.1:3308)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移模式，自动创建或修改数据库表结构
	db.AutoMigrate(&Student{}, &Course{}, &Score{})

	// importData()

	Query()
}

func importData() {
	// 插入学生数据
	students := []Student{
		{Name: "Alice", Age: 20, Class: "Class 1"},
		{Name: "Bob", Age: 21, Class: "Class 2"},
		{Name: "Charlie", Age: 19, Class: "Class 1"},
		{Name: "David", Age: 22, Class: "Class 2"},
		{Name: "Eve", Age: 20, Class: "Class 1"},
	}
	db.Create(&students)

	// 插入课程数据
	courses := []Course{
		{Name: "Math", Teacher: "Mr. A", Room: "101"},
		{Name: "Physics", Teacher: "Mr. B", Room: "102"},
		{Name: "Chemistry", Teacher: "Ms. C", Room: "103"},
		{Name: "Biology", Teacher: "Mr. D", Room: "104"},
		{Name: "Computer Science", Teacher: "Ms. E", Room: "105"},
	}
	db.Create(&courses)

	// 插入成绩数据
	scores := []Score{
		{StudentID: 1, CourseID: 1, Score: 85},
		{StudentID: 1, CourseID: 2, Score: 90},
		{StudentID: 2, CourseID: 1, Score: 75},
		{StudentID: 2, CourseID: 3, Score: 80},
		{StudentID: 3, CourseID: 4, Score: 95},
		{StudentID: 3, CourseID: 5, Score: 88},
		{StudentID: 4, CourseID: 2, Score: 78},
		{StudentID: 4, CourseID: 3, Score: 82},
		{StudentID: 5, CourseID: 4, Score: 92},
		{StudentID: 5, CourseID: 5, Score: 89},
	}
	db.Create(&scores)
}
