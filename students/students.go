package students

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

type Profile struct {
	First    string
	Last     string
	Id       string
	Email    string
	CourseId string
}

// List returns a list of all the students in a course with a specific Id
func List(client *http.Client, id string) []Profile {

	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}

	loop := true
	// My attempt to get courseWork using courseId 126909787383
	pt := "nil"
	roster := []Profile{}
	for loop == true {

		// if there is a page token, you should not return yet
		if pt != "nil" {
			students, err := srv.Courses.Students.List(id).PageToken(pt).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve students. %v", err)
			}
			//roster := []Profile{}
			if len(students.Students) > 0 {

				if students.NextPageToken != "" {
					pt = students.NextPageToken
				} else {
					pt = "nil"
					for _, s := range students.Students {
						studentProfile := Profile{
							First:    s.Profile.Name.GivenName,
							Last:     s.Profile.Name.FamilyName,
							Id:       s.Profile.Id,
							Email:    s.Profile.EmailAddress,
							CourseId: s.CourseId,
						}
						roster = append(roster, studentProfile)
					}
					break
				}
				for _, s := range students.Students {
					studentProfile := Profile{
						First:    s.Profile.Name.GivenName,
						Last:     s.Profile.Name.FamilyName,
						Id:       s.Profile.Id,
						Email:    s.Profile.EmailAddress,
						CourseId: s.CourseId,
					}
					roster = append(roster, studentProfile)
				}
			} else {
				fmt.Print("No students found.")
			}
			/* return roster */
		} else {

			students, err := srv.Courses.Students.List(id).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve students. %v", err)
			}
			/* roster := []Profile{} */
			if len(students.Students) > 0 {
				fmt.Print("\nStudents:\n")

				if students.NextPageToken != "" {
					pt = students.NextPageToken
				} else {
					pt = "nil"
					for _, s := range students.Students {
						studentProfile := Profile{
							First:    s.Profile.Name.GivenName,
							Last:     s.Profile.Name.FamilyName,
							Id:       s.Profile.Id,
							Email:    s.Profile.EmailAddress,
							CourseId: s.CourseId,
						}
						roster = append(roster, studentProfile)
					}
					break
				}

				for _, s := range students.Students {
					studentProfile := Profile{
						First:    s.Profile.Name.GivenName,
						Last:     s.Profile.Name.FamilyName,
						Id:       s.Profile.Id,
						Email:    s.Profile.EmailAddress,
						CourseId: s.CourseId,
					}
					roster = append(roster, studentProfile)
				}
			} else {
				fmt.Print("No students found.")
			}
			/* return roster */
		}
	}
	fmt.Printf("Total: %d\n", len(roster))
	return roster
}

/*     students, err := srv.Courses.Students.List(id).Do()
	 *     if err != nil {
	 *         log.Fatalf("Unable to retrieve students. %v", err)
	 *     }
	 *     roster := []Profile{}
	 *     if len(students.Students) > 0 {
	 *         //fmt.Print("\nStudents:\n")
	 *
	 P>
 *         for _, s := range students.Students {
 *             studentProfile := Profile{
 *                 First:    s.Profile.Name.GivenName,
 *                 Last:     s.Profile.Name.FamilyName,
 *                 Id:       s.Profile.Id,
 *                 Email:    s.Profile.EmailAddress,
 *                 CourseId: s.CourseId,
 *             }
 *             roster = append(roster, studentProfile)
 *         }
 *     } else {
 *         fmt.Print("No students found.")
 *     }
 *     return roster */
