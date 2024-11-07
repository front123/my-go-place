package models

import (
	"encoding/json"
	"fmt"
)

type Dto struct {
	StudentID   uint
	StudentName string
	Score       int

	CourseDetailStr string
	CourseInfo      *CourseSimple `gorm:"embedded;embeddedPrefix:course_"`
}

type CourseSimple struct {
	Name    string
	Teacher string
	Room    string
}

func Query() {
	var ret []Dto
	db.Table("students").
		Joins("left join scores on students.id = scores.student_id").
		Joins("left join courses on scores.course_id = courses.id").
		Select(
			"students.id as student_id",
			"students.name as student_name",
			"courses.name as course_name",
			"courses.teacher as course_teacher",
			"courses.room as course_room",
			"scores.score as score",
			"concat_ws(',', courses.name, courses.teacher, courses.room) as course_detail_str",
		).Scan(&ret)

	if len(ret) > 0 {
		bs, _ := json.Marshal(ret)
		fmt.Println(string(bs))
	}
}
