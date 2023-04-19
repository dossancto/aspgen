package generate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lu-css/aspgen/src/general"
	"github.com/manifoldco/promptui"
)

func Generate() {

	dbProviders := []string{
		"area",
		"controller",
		"identity",
		"minimalapi",
		"razorpage",
		"view",
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "\U0001F336 {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose what you want to generate",
		Items:     dbProviders,
		Templates: templates,
		Size:      8,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		general.Exit()
		return
	}

	switch dbProviders[i] {
	case "controller":
		genController()
	}
}

func genController() {
	controllerName := getControllerName()
	DbProvider := getDbProvider()
	useDefaultLayout := getUseDefaultLayout()
	model := getModel()
	dbContext := getDbContext()

	controller := AspController{
		ControllerName:   controllerName,
		Model:            model,
		UseDefaultLayout: useDefaultLayout,
		DbContext:        dbContext,
		DbProvider:       DbProvider,
	}

	controller.runAspnetCodegenerator()
}

func getDbContext() string {
	prompt := promptui.Prompt{
		Label: "Db Context",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result

}

func getControllerName() string {
	prompt := promptui.Prompt{
		Label:    "Controller Name",
		Validate: NonBlankInput,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func getUseDefaultLayout() bool {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Blank Text")
		}

		lowerInput := strings.ToLower(input)

		if lowerInput != "y" && lowerInput != "n" {
			return errors.New("Place use 'y' or 'n'")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Use Default Layout?",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	return strings.ToLower(result) == "y"

}

func getDbProvider() string {
	dbProviders := []string{
		"Postgres",
		"SQLserver",
		"Cosmos",
		"Sqlite",
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "Database Provider: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a Model",
		Items:     dbProviders,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return strings.ToLower(dbProviders[i])

}

func getModel() string {
	models, err := GetAllModels("Models")

	if err != nil {
		fmt.Printf("Somwthing went wrong\n%v\n", err)
		general.Exit()
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }} ({{ . | red }})",
		Inactive: "  {{ . | cyan }} ({{ .| red }})",
		Selected: " Selected Model: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a Model",
		Items:     models,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		general.Exit()
		return ""
	}

	return models[i]

}
