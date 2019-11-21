package story

import (
	. "../parser"
	"encoding/json"
	"fmt"
	"github.com/teris-io/shortid"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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

var storyTemplate = `# id: {{.Id}}
# startDate: {{.StartDate}}
# endDate: {{.EndDate}}
# priority: {{.Priority}}
# status: {{.Status}}
# author: {{.Author}}
# language: zh-CN
@math
功能:{{.Title}}

  背景:
    你好啊®

  场景: 大纲
    假设: 我需要
    当: 我blbla
    并且: 这很不错
    那么: 就好了
`

func SyncStory() {

}

func ListStory() []StoryModel {
	stories := []StoryModel{}
	//for _, f := range storyList {
	//	story := StoryModel{}
	//	if err := json.Unmarshal([]byte(f), &story); err != nil {
	//		fmt.Println("Error", err)
	//	}
	//	stories = append(stories, story)
	//}
	//
	return stories
}

func CreateStory(content string) {
	storyId := buildStoryId()

	date := &StoryDate{getNow(), getNow()}
	story := &StoryModel{storyId, content, "", "", "", *date, "", "", "", "", ""}

	filePath := getFilePath(storyId, content)
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

func getFilePath(u1 string, content string) string {
	fileName := u1 + "-" + updateFileName(content)
	filePath := "stories/" + fileName + ".feature"
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

	fileName := getFeatureFileNameById(id)
	if fileName == "" {
		log.Fatal("error id, %s", id)
	}

	filePath := "stories/" + fileName
	story = parseFeature(filePath)
	story.Author = userName

	file, e := os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	if e != nil {
		log.Fatal("open file error %s:", e)
		return
	}
	saveStory(file, &story)
}

func ChangeStoryStatus(id string, status string) {
	var story StoryModel

	fileName := getFeatureFileNameById(id)
	if fileName == "" {
		log.Fatal("error id, %s", id)
	}

	filePath := "stories/" + fileName
	story = parseFeature(filePath)
	story.Status = status
	story.EndDate = getNow()

	file, e := os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	if e != nil {
		log.Fatal("open file error %s:", e)
		return
	}
	saveStory(file, &story)
}

func getNow() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func getFeatureFileNameById(id string) string {
	files, err := ioutil.ReadDir("stories/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), id)
		if strings.Contains(f.Name(), id) {
			return f.Name()
		}
	}

	return ""
}

func parseFeature(path string) StoryModel {
	var storyModel StoryModel
	app := NewFeatureApp()
	results := app.Start(path)
	jsonStr, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	storyModel.StartDate = results["startDate"]
	storyModel.EndDate = results["endDate"]

	_ = json.Unmarshal(jsonStr, &storyModel)
	return storyModel
}
