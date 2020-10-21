package main

import (
	"fmt"

	/* co "github.com/jdvober/goClassroomTools/courses" */
	co "github.com/jdvober/goClassroomTools/courses"
	"github.com/jdvober/goClassroomTools/students"
	stu "github.com/jdvober/goClassroomTools/students"
	auth "github.com/jdvober/goGoogleAuth"
	sh "github.com/jdvober/goSheets/values"
)

func main() {
	client := auth.Authorize()

	spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
	rangeData := "sheet2!A2"

	courses := co.List(client)
	var studentProfiles []students.Profile

	for _, course := range courses {
		studentList := stu.List(client, course.Id) // CourseId Email Id First Last
		for _, student := range studentList {

			studentProfiles = append(studentProfiles, student)
		}
	}

	// Get all courses
	values := make([][]interface{}, len(studentProfiles))
	counter := 0
	for _, course := range courses {
		students := stu.List(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)

			values[counter] = []interface{}{s.First, s.Last, s.Email, course.Name, course.Id}
			counter++
		}
	}

	sh.BatchUpdate(client, spreadsheetId, rangeData, values)
	fmt.Println("Finished main()")
}
