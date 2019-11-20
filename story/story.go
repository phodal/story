package story

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type StoryDate struct {
	StartDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
}

type StoryModel struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	AC           string `json:"ac,omitempty"`
	Condition    string `json:"condition,omitempty"`
	StoryDate    `json:"date,omitempty"`
	Dependencies string `json:"dependencies,omitempty"`
	Priority     string `json:"priority,omitempty"`
}

var driver *Driver

func InitStory() {
	driver, _ = NewZhu("story/db")
}

func AddStory() {

}

func CreateStory(content string) {
	u1 := uuid.Must(uuid.NewUUID())

	date := &StoryDate{time.Now(), time.Now()}
	story := &StoryModel{u1.String(), content, "", "", "", *date, "", ""}

	error := driver.Write("story", story.Id, story)
	if error != nil {
		log.Fatal(error)
	}
}
