package main

import (
	//"./cmd"
	. "./parser"
	. "./story"
	"encoding/json"
	"log"
)

var storyModel StoryModel

func main() {
	//cmd.Execute()

	app := NewFeatureApp()
	results := app.Start("tests/test.feature")
	jsonStr, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(jsonStr, &storyModel)
	log.Println(storyModel)
}