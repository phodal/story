package story

import (
	. "../parser"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func GetFeatureFileNameById(id string) string {
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

func ParseFeature(path string) StoryModel {
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
