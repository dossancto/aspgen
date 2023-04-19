package main

import (
	"fmt"

	"github.com/lu-css/aspgen/src/generate"
	"github.com/lu-css/aspgen/src/validations"
)

func main() {
	inCsprojDir := validations.ExistsCsProj()

	if !inCsprojDir {
		fmt.Println("Csproj not found.")
		return
	}

	generate.Generate()
}
