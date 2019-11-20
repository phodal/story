package story

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"text/template"
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

var storyTemplate = `作为

我想要{{.Title}}

这样就能
`

func InitStory() {
	driver, _ = NewZhu("stories/db")
}

func ListStory() {
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

	t, _ := template.New("story").Parse(storyTemplate)
	file, err := os.Create("stories/" + u1 + ".md")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(file, &story)
	if err != nil {
		log.Println("template file: ", err)
		return
	}

	err = driver.Write("stories", story.Id, story)
	if err != nil {
		log.Fatal(err)
	}
}
