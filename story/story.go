package story

import (
	"encoding/json"
	"fmt"
	"github.com/teris-io/shortid"
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
	Author       string `json:"author,omitempty"`
	Status       string `json:"status,omitempty"`
}

var driver *Driver

var storyTemplate = `作为

我想要{{.Title}}

这样就能
`

func InitStory() {
	driver, _ = NewZhu("stories/db")
}

func ListStory() []StoryModel {
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
		stories = append(stories, story)
	}

	return stories
}

func CreateStory(content string) {
	u1 := buildStoryId()

	date := &StoryDate{time.Now(), time.Now()}
	story := &StoryModel{u1, content, "", "", "", *date, "", "", "", ""}

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

func buildStoryId() string {
	str, _ := shortid.Generate()
	return str
}

func PickStory(id string, userName string) {
	var story StoryModel
	err := driver.Read("stories", id, &story)
	if err != nil {
		log.Fatal(err)
	}

	story.Author = userName

	err = driver.Write("stories", id, &story)
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeStoryStatus(id string, status string) {
	var story StoryModel
	err := driver.Read("stories", id, &story)
	if err != nil {
		log.Fatal(err)
	}

	story.Status = status

	err = driver.Write("stories", id, &story)
	if err != nil {
		log.Fatal(err)
	}
}