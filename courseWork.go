package gclass

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

type Date struct {
	Year  int64
	Month int64
	Day   int64
}
type TimeOfDay struct {
	Hours   int64
	Minutes int64
}

type Coursework struct {
	CourseID      string
	Id            string
	Title         string
	Description   string
	State         string
	CreationTime  string
	UpdateTime    string
	DueDate       Date
	DueTime       TimeOfDay
	ScheduledTime string
	MaxPoints     float64
	TopicID       string
}

// List returns a list of all the coursework in a course with a specific Id
func ListCoursework(client *http.Client, id string) []Coursework {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// My attempt to get courseWork using courseId 126909787383
	r, err := srv.Courses.CourseWork.List(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve coursework. %v", err)
	}
	courseworkList := []Coursework{}
	if len(r.CourseWork) > 0 {
		//fmt.Print("\nCoursework:\n")
		for _, cw := range r.CourseWork {

			cwork := Coursework{

				CourseID:      cw.CourseId,
				Id:            cw.Id,
				Title:         cw.Title,
				Description:   cw.Description,
				State:         cw.State,
				CreationTime:  cw.CreationTime,
				UpdateTime:    cw.UpdateTime,
				DueDate:       Date{Year: cw.DueDate.Year, Month: cw.DueDate.Month, Day: cw.DueDate.Day},
				DueTime:       TimeOfDay{Hours: cw.DueTime.Hours, Minutes: cw.DueTime.Minutes},
				ScheduledTime: cw.ScheduledTime,
				MaxPoints:     cw.MaxPoints,
				TopicID:       cw.TopicId,
			}
			courseworkList = append(courseworkList, cwork)
		}

	} else {
		fmt.Print("No coursework found.")
	}
	return courseworkList
}

/*JSON Response Look Like This:

{
"courseWork": [
{
	"courseId": "126909787383",
	"id": "178503269396",
	"title": "7.4.4 - Discuss",
	"description": "Please take part in a discussion addressing the questions posed in 7.4.4 on APEX, ...etc",
	"state": "PUBLISHED",
	"alternateLink": "https://classroom.google.com/c/MTI2OTA5Nzg3Mzgz/sa/MTc4NTAzMjY5Mzk2/details",
	"creationTime": "2020-09-29T14:05:12.405Z",
	"updateTime": "2020-09-29T14:06:59.230Z",
	"dueDate": {
	"year": 2020,
	"month": 10,
	"day": 16
	},
	"dueTime": {
	"hours": 3,
	"minutes": 59
	},
	"maxPoints": 15,
	"workType": "SHORT_ANSWER_QUESTION",
	"submissionModificationMode": "MODIFIABLE",
	"assigneeMode": "ALL_STUDENTS",
	"creatorUserId": "118026622292638242171",
	"topicId": "178468686835"
},

*/
