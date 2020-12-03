package main

import (
	"fmt"

	"github.com/jdvober/gauth"
	"github.com/jdvober/gclass"
)

func main() {
	client := gauth.Authorize()

	// Get all courses
	courses := gclass.ListCourses(client)
	// For each course,list all students in the class in a different column on google sheets
	for _, course := range courses {
		fmt.Printf("\nGetting students for %s\n", course.Name)
		students := st.List(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)
		}
	}
}
