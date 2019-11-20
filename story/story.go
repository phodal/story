package story

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
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
	driver, _ = NewZhu("stories/db")
}

func ListStory()  {
	storyList, error := driver.ReadAll("stories")
	if error != nil {
		log.Fatal(error)
	}

	stories := []StoryModel{}
	for _, f := range storyList {
		story := StoryModel{}
		if err := json.Unmarshal([]byte(f), &story); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(story.Title)
		stories = append(stories, story)
	}
}

func CreateStory(content string) {
	u1 := uuid.Must(uuid.NewUUID()).String()

	date := &StoryDate{time.Now(), time.Now()}
	story := &StoryModel{u1, content, "", "", "", *date, "", ""}

	fileError := ioutil.WriteFile("stories/"+u1+".md", []byte(`作为
我想要
这样就能
`), 0644)
	if fileError != nil {
		log.Fatal(fileError)
	}

	error := driver.Write("stories", story.Id, story)
	if error != nil {
		log.Fatal(error)
	}
}
