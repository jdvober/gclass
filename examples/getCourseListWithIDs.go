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
	for _, course := range courses {
		fmt.Printf("%s: %s\n", course.Name, course.Id)
	}
}
