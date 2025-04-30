import data
import (
	"time"
)



type Work struct {
	ID int64; //id of work
	CrearedAt time.Time; //when posted
	LastEditedAt time.Time; //when last edit
	Title string; //Title of work
	Author string; //username of the author
	FileName string; //link to the file
	WorkType string; //type of work: Object or Environment
}


