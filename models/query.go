package models

import (
	"encoding/json"
	"fmt"
)

type Dto struct {
	StudentID uint
	StudentName string
	Course string
	Score int

	CourseDetail string
}

func Query() {
	var ret []*Dto
	db.Table("students").
	Joins("inner join scores on students.id = scores.student_id").
	Joins("inner join courses on scores.course_id = courses.id").
	Select(
		"students.id as student_id",
		"students.name as student_name",
		"courses.name as course",
		"scores.score as score",
		"concat_ws(',', courses.name, courses.teacher, courses.room) as course_detail",
	).Scan(&ret)

	if len(ret) > 0 {
		bs, _ := json.Marshal(ret)
		fmt.Println(string(bs))
	}
}