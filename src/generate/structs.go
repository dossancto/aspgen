package generate

import (
	"fmt"
	"github.com/lu-css/aspgen/src/general"
	"os/exec"
)

type AspController struct {
	ControllerName   string
	Model            string
	UseDefaultLayout bool
	DbContext        string
	DbProvider       string
}

func (controller AspController) runAspnetCodegenerator() {
	var udl string

	if controller.UseDefaultLayout {
		udl = "-udl"
	}
	cmd := exec.Command("dotnet", "aspnet-codegenerator", "controller",
		"-name", controller.ControllerName,
		"-m", controller.Model,
		"-dbProvider", controller.DbProvider,
		"-dc", controller.DbContext,
		udl)

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", out)
		general.Exit()
	}

	fmt.Printf("%s", out)

	if err != nil {
		fmt.Printf("%s\n", err)
		general.Exit()
	}

}
