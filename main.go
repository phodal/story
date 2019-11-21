package main

import (
	//"./cmd"
	. "./parser"
)

func main() {
	//cmd.Execute()

	app := NewFeatureApp()
	app.Start("tests/test.feature")
}