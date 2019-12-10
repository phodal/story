package story

import (
	"log"
	"os"
	. "story/parser"
	"time"
)

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
	SyncFeatures()
}

func ListStory() []StoryModel {
	return GetFeaturesByPath()
}

func CreateStory(content string) {
	storyId := GenerateId()

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
	SaveStoryToFile(file, story)
	defer file.Close()
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

