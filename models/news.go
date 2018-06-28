package models 

type NewsPost struct {
        ID int
        Title string
        Body string
        Date string
        Author string
}

type NewsComment struct {
	ID int
	PID int // parent = newspost
	Body string
	Date string
	Author string
}
