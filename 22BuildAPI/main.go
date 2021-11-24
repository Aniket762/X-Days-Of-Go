package main

import "fmt"

// model for course
type Course struct {
	CourseId  string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// middlewares
func (c *Course ) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// dummy database
var courses [] Course

func main()  {
	fmt.Println("Getting into building API 🙏")


}