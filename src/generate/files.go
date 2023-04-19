package generate

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetAllModels(dirName string) ([]string, error) {
	files, err := os.ReadDir(dirName)

	var models []string

	if err != nil {
		println("error")
		fmt.Println(err)
	}

	for _, file := range files {
		fileName := file.Name()

		models = append(models, strings.Replace(fileName, ".cs", "", -1))
	}

	if len(models) > 0 {
		return models, nil
	}

	return models, errors.New("none Models")
}
