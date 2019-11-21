package story

import (
	"bytes"
	"crypto/md5"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return value, nil
}

func Md5SumFolder(folder string) (map[string][md5.Size]byte, error) {
	results := make(map[string][md5.Size]byte)
	_ = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//判断文件属性
		if info.Mode().IsRegular() {
			result, err := Md5SumFile(path)
			if err != nil {
				return err
			}
			results[path] = result
		}
		return nil
	})
	return results, nil
}

func GetFilePath(id string, content string) string {
	fileName := BuildFileName(id, content)
	filePath := "stories/" + fileName
	return filePath
}

func BuildFileName(id string, content string) string {
	return id + "-" + UpdateFileName(content) + ".feature"
}

func UpdateFileName(name string) string {
	rex := regexp.MustCompile(`[，,。 ！？]`)
	return rex.ReplaceAllString(name, "")
}

func UpdateStoryByFilePath(filePath string, model *StoryModel) {
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

	WriteNewTemplateToFile(filePath, result)
}

func WriteNewTemplateToFile(filePath string, result string) {
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

func SaveStoryToFile(file *os.File, story *StoryModel) {
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
