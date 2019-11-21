package story

import (
	. "../parser"
	"github.com/teris-io/shortid"
	"log"
	"os"
	"text/template"
	"time"
)

type StoryDate struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
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

var infoTemplate = `# id: {{.Id}}
# startDate: {{.StartDate}}
# endDate: {{.EndDate}}
# priority: {{.Priority}}
# status: {{.Status}}
# author: {{.Author}}
# title: {{.Title}}
# language: zh-CN`

var storyTemplate = infoTemplate + `
@math
功能:{{.Title}}

  场景:
    假设:
    当:
    并且:
    那么:
`

var commentPosition = CommentPosition{}

func SyncStory() {

}

func ListStory() []StoryModel {
	stories := []StoryModel{}

	return stories
}

func CreateStory(content string) {
	storyId := buildStoryId()

	date := &StoryDate{getNow(), getNow()}
	story := &StoryModel{storyId, content, "", "", "", *date, "", "", "", "", ""}

	filePath := GetFilePath(storyId, content)
	file, err := os.Create(filePath)
	hashValue, err := Md5SumFile(filePath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	story.Hash = string(hashValue[:])
	saveStory(file, story)
}

func saveStory(file *os.File, story *StoryModel) {
	t, err := template.New("story").Parse(storyTemplate)
	if err != nil {
		log.Println("template error: ", err)
		return
	}
	err = t.Execute(file, &story)
	if err != nil {
		log.Println("template file: ", err)
		return
	}
}

func buildStoryId() string {
	str, _ := shortid.Generate()
	return str
}

func PickStory(id string, userName string) {
	var story StoryModel

	fileName := GetFeatureFileNameById(id)
	if fileName == "" {
		log.Fatal("error id, %s", id)
	}

	filePath := "stories/" + fileName
	story = ParseFeature(filePath)
	story.Author = userName

	UpdateStoryByFilePath(filePath, &story)
}

func ChangeStoryStatus(id string, status string) {
	var story StoryModel

	fileName := GetFeatureFileNameById(id)
	if fileName == "" {
		log.Fatal("error id, %s", id)
	}

	filePath := "stories/" + fileName
	story = ParseFeature(filePath)
	story.Status = status
	story.EndDate = getNow()

	UpdateStoryByFilePath(filePath, &story)
}

func getNow() string {
	return time.Now().UTC().Format(time.RFC3339)
}

