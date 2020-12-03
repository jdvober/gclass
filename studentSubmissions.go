package gclass

import (
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

// List returns a list of all the students in a course with a specific Id
func ListStudentSubmissions(client *http.Client, id string, cwid string) []*classroom.StudentSubmission {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}

	courseID := id
	courseWorkID := cwid
	// pToken := ""

	Submissions := []*classroom.StudentSubmission{}

	// My attempt to get studentSubmissions using courseId 126909787383 and courseWorkId "-" which returns all studentSubmissions

	// https://godoc.org/google.golang.org/api/classroom/v1#CoursesCourseWorkStudentSubmissionsService
	res, err := srv.Courses.CourseWork.StudentSubmissions.List(courseID, courseWorkID).PageSize(2).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve submission. %v", err)
	}
	for _, s := range res.StudentSubmissions {
		Submissions = append(Submissions, s)
	}

	return Submissions
}
