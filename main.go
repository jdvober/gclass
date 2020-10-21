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
	for _, course := range courses {
		fmt.Printf("%s: %s\n", course.Name, course.Id)
	}
	// For each course,list all students in the class in a different column on google sheets
	for _, course := range courses {
		fmt.Printf("\nGetting students for %s\n", course.Name)
		students := st.List(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)
		}
	}

	// submissions := subs.List(client, "126909787383", "-")

	/* courses := co.List(client) */
	// courseWork := cw.List(client, "126909787383")

	// fmt.Println("\nCourses (Course ID): >>")
	// for cl := range courseList {
	// 	spew.Dump(courses[cl])
	// }

	// fmt.Println("\nCoursework: >>")
	// for c := range courseWork {
	// 	spew.Dump(courseWork[c])
	// }

	// fmt.Println("\nStudents: >>")
	// for s := range students {
	// 	spew.Dump(students[s])
	// }

	/* fmt.Println("\nSubmissions: >>")
	 * for s := range submissions {
	 *     // fmt.Printf("\nSubmission %d\n%+v\n", s, submissions[s])
	 *     spew.Dump(submissions[s])
	 * } */

	// fmt.Println("\n<<Press any key to exit...>>")
	// fmt.Scanln()

}
