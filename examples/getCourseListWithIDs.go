package main

import (
	"fmt"

	auth "github.com/jdvober/goGoogleAuth"

	/* subs "github.com/jdvober/goClassroomTools/studentSubmissions" */

	co "github.com/jdvober/goClassroomTools/courses"
	/* cw "github.com/jdvober/goClassroomTools/courseWork" */)

func main() {
	client := auth.Authorize()

	// Get all courses
	courses := co.List(client)
	for _, course := range courses {
		fmt.Printf("%s: %s\n", course.Name, course.Id)
	}
}
