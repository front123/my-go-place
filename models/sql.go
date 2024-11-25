package models

import "fmt"

func SqlAssign() {

	// FirstOrCreate find first
	// if exist, updated model with input arguments => stu1 load model from db
	// if not exist, create model with input argument and model origin field values => stu1 use origin
	// stu1 := &Student{Age: 111}
	// db.Where(&Student{Name: "Bob"}).Assign(map[string]interface{}{
	// 	"Class": "100Class",
	// }).FirstOrCreate(stu1)
	// fmt.Printf("%+v\n", stu1)

	// FirstOrInit find first one match
	// if exist, load model to struct
	// if not exist, update struct with input arguments
	// never modify database
	stu2 := &Student{Age: 122}
	db.Where(&Student{Name: "Bob2"}).Assign(map[string]interface{}{
		"Class": "101Class",
	}).FirstOrInit(stu2)
	fmt.Printf("%+v\n", stu2)
}
