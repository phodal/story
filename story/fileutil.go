package story

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"path/filepath"
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