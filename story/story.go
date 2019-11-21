package story

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/teris-io/shortid"
	"hash"
	"log"
	"os"
	"regexp"
	"text/template"
	"time"
)

type StoryDate struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
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
	Hash         string
}

var driver *Driver
var fileMd5 hash.Hash

var storyTemplate = `作为

我想要{{.Title}}

这样就能
`

func InitStory() {
	driver, _ = NewZhu("stories/db")
	fileMd5 = md5.New()

	SyncStory()
}

func SyncStory() {

}

func ListStory() []StoryModel {
	storyList, err := driver.ReadAll("stories")
	if err != nil {
		log.Fatal(err)
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
	storyId := buildStoryId()

	date := &StoryDate{time.Now(), time.Now()}
	story := &StoryModel{storyId, content, "", "", "", *date, "", "", "", "", ""}

	t, _ := template.New("story").Parse(storyTemplate)
	filePath := getFilePath(storyId, content)
	file, err := os.Create(filePath)
	hashValue, err := Md5SumFile(filePath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	story.Hash = string(hashValue[:])
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

func getFilePath(u1 string, content string) string {
	fileName := u1 + "-" + updateFileName(content)
	filePath := "stories/docs/" + fileName + ".md"
	return filePath
}

func updateFileName(name string) string {
	rex := regexp.MustCompile(`[，,。 ！？]`)
	return rex.ReplaceAllString(name, "")
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
	if status == "done" {
		story.EndDate = time.Now()
	}

	err = driver.Write("stories", id, &story)
	if err != nil {
		log.Fatal(err)
	}
}