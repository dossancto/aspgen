package validations

import (
	"fmt"
	"os"
	"strings"
)

func ExistsCsProj() bool {
	dirName, err := os.Getwd()

	if err != nil {
		println("error")
		fmt.Println(err)
	}

	files, err := os.ReadDir(dirName)

	if err != nil {
		println("error")
		fmt.Println(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if exists := strings.Contains(fileName, ".csproj"); exists {
			return true
		}
	}

	return false
}
