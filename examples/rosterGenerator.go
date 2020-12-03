package main

import (
	"fmt"

	"github.com/jdvober/gauth"
	"github.com/jdvober/gclass"
	"github.com/jdvober/gsheets"
)

func main() {
	client := gauth.Authorize()

	spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
	rangeData := "sheet2!A2"

	courses := gclass.ListCourses(client)
	var studentProfiles []students.Profile

	for _, course := range courses {
		studentList := gclass.ListStudents(client, course.Id) // CourseId Email Id First Last
		for _, student := range studentList {

			studentProfiles = append(studentProfiles, student)
		}
	}

	// Get all courses
	values := make([][]interface{}, len(studentProfiles))
	counter := 0
	for _, course := range courses {
		students := gclass.ListStudents(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)

			values[counter] = []interface{}{s.First, s.Last, s.Email, course.Name, course.Id}
			counter++
		}
	}

	gsheets.BatchUpdateValues(client, spreadsheetId, rangeData, values)
	fmt.Println("Finished main()")
}
