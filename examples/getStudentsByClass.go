package main

import (
	"fmt"

	auth "github.com/jdvober/goGoogleAuth"

	/* subs "github.com/jdvober/goClassroomTools/studentSubmissions" */
	st "github.com/jdvober/goClassroomTools/students"

	co "github.com/jdvober/goClassroomTools/courses"
	/* cw "github.com/jdvober/goClassroomTools/courseWork" */)

func main() {
	client := auth.Authorize()

	// Get all courses
	courses := co.List(client)
	// For each course,list all students in the class in a different column on google sheets
	for _, course := range courses {
		fmt.Printf("\nGetting students for %s\n", course.Name)
		students := st.List(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)
		}
	}

}
