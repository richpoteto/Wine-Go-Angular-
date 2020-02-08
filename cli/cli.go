package cli

import (
	"fmt"
	"github.com/shpota/goxygen/codegen"
	"regexp"
)

func Start(commands []string) {
	if len(commands) == 1 && commands[0] == "help" {
		fmt.Println(usage)
		return
	}
	if len(commands) == 2 && commands[0] == "init" {
		projectName := commands[1]
		validName, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", projectName)
		if validName {
			codegen.Generate(projectName)
			return
		}
		fmt.Println(invalidName)
		return
	}
	fmt.Println(wrongInput)
}

const usage = `Usage:

   go run github.com/shpota/goxygen init <project-name>

Generates a web project with the following structure:
server side Go code, webapp with corresponding React
components, a Dockerfile that builds the application,
and Docker Compose scripts to run the app in dev and
prod modes.`

const wrongInput = `Wrong input! Use:
   go run github.com/shpota/goxygen init <project-name>
In order to get more details run:
   go run github.com/shpota/goxygen help`

const invalidName = `Project name is not valid. The allowed symbols are letters,
numbers, underscores, and dashes.`
