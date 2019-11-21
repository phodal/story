package story

import (
	. "../parser"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var storiesPath = "stories/"

func GetFeatureFileNameById(id string) string {
	files, err := ioutil.ReadDir(storiesPath)
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

func GetFeaturesByPath() []StoryModel {
	files, err := ioutil.ReadDir(storiesPath)
	if err != nil {
		log.Fatal(err)
	}

	var stories []StoryModel
	for _, f := range files {
		feature := ParseFeature(storiesPath + f.Name())
		stories = append(stories, feature)
	}

	return stories
}

func SyncFeatures() {
	files, err := ioutil.ReadDir(storiesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		feature := ParseFeature(storiesPath + f.Name())

		reallyFileName := BuildFileName(feature.Id, feature.Title)
		if f.Name() != reallyFileName {
			err := os.Rename(storiesPath+f.Name(), storiesPath+reallyFileName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
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
