package generate

import (
	"fmt"
	"os"

	"github.com/lu-css/aspgen/src/utils"
	"github.com/paulrademacher/climenu"
)

func Generate() {
	menu := climenu.NewButtonMenu("Aspgen", "Choose what you want to generate")

	menu.AddMenuItem("area", "area")
	menu.AddMenuItem("controller", "controller")
	menu.AddMenuItem("identity", "identity")
	menu.AddMenuItem("minimalapi", "minimalapi")
	menu.AddMenuItem("razorpage", "razorpage")
	menu.AddMenuItem("view", "view")

	action, escaped := menu.Run()

	if escaped {
		os.Exit(0)
	}

	switch action {
	case "area":
		Area()
	case "controller":
		genController()
	}
}

func Area() {

}

func genController() {
	controllerName := climenu.GetText("ControllerName", "nothing")

	checkboxController := climenu.NewCheckboxMenu("Some Options", "Select options (using space key)", "OK", "CANCEL")
	checkboxController.AddMenuItem("useAsyncActions", "useAsyncActions")
	checkboxController.AddMenuItem("useAsyncActions", "useAsyncActions")

	dbProviderMenu := climenu.NewButtonMenu("Database Provider", "The database to be used")
	dbProviderMenu.AddMenuItem("Postgres", "postgres")
	dbProviderMenu.AddMenuItem("SQLserver", "sqlserver")
	dbProviderMenu.AddMenuItem("cosmos", "cosmos")
	dbProviderMenu.AddMenuItem("sqlite", "sqlite")

	dbProvider, escaped := dbProviderMenu.Run()

	if escaped {
		os.Exit(0)
	}

	useDefaultLayout := utils.TrueOrFalse("Use Default Layout?", true)

	modelsMenu := climenu.NewButtonMenu("Models", "Select a model")

	models, err := GetAllModels("Models")

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, model := range models {
		modelsMenu.AddMenuItem(model, model)
	}

	choosedModel, escaped := modelsMenu.Run()

	if escaped {
		os.Exit(0)
	}

	dbContext := climenu.GetText("DbContext Path", "ApplicationDbContext")

	println(controllerName)
	println(dbProvider)
	println(useDefaultLayout)
	println(choosedModel)
	println(dbContext)
}
