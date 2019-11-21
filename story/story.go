package story

import (
	. "../parser"
	"bytes"
	"encoding/json"
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

var infoTemplate = `# id: {{.Id}}
# startDate: {{.StartDate}}
# endDate: {{.EndDate}}
# priority: {{.Priority}}
# status: {{.Status}}
# author: {{.Author}}
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

	updateStory(filePath, &story)
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

	updateStory(filePath, &story)
}

func updateStory(filePath string, model *StoryModel) {
	tpl, err := template.New("info").Parse(infoTemplate)
	if err != nil {
		log.Fatal(err)
		return
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, model); err != nil {
		log.Fatal(err)
		return
	}

	result := buffer.String()

	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	newLines := make([]string, len(lines)-commentPosition.End)

	copy(newLines, lines[commentPosition.End:])

	output := result + "\n" + strings.Join(newLines, "\n")
	err = ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
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
		if strings.Contains(f.Name(), id) {
			return f.Name()
		}
	}

	return ""
}

func parseFeature(path string) StoryModel {
	var storyModel StoryModel
	app := NewFeatureApp()
	commentStruct := app.Start(path)
	results := commentStruct.CommentsMap
	commentPosition = commentStruct.CommentPosition
	jsonStr, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	storyModel.StartDate = results["startDate"]
	storyModel.EndDate = results["endDate"]

	_ = json.Unmarshal(jsonStr, &storyModel)
	return storyModel
}
