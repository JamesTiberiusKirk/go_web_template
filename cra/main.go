package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/JamesTiberiusKirk/go_web_template/cra/cra"
)

// var (
// 	exampleOptions = cra.Options{
// 		ProjectParentDir: "/home/darthvader/Projects",
// 		// ProjectParentDir:    "/Users/darthvader/Projects",
// 		TemplateDir:         ".",
// 		ProjectName:         "test_project",
// 		GoProjectModuleName: "example.com/me/test_project",
// 		Vendoring:           false,
// 		Selections: cra.Selections{
// 			API:  true,
// 			Site: true,
// 			SiteConfig: &cra.Site{
// 				Templating: true,
// 				SSR:        true,
// 				Static:     true,
// 				SPA:        true,
// 			},
// 			Auth: cra.Session,
// 		},
// 	}
// )

func main() {

	flagOptions := buildFlags()

	if flagOptions.ProjectDir == "" ||
		flagOptions.GoModuleName == "" ||
		flagOptions.ProjectName == "" {
		fmt.Println("Missing params")
		flag.PrintDefaults()
		os.Exit(1)
	}

	projectOptions := cra.Options{
		ProjectParentDir:    flagOptions.ProjectDir,
		ProjectName:         flagOptions.ProjectName,
		GoProjectModuleName: flagOptions.GoModuleName,
		Vendoring:           flagOptions.Vendoring,
		Selections: cra.Selections{
			API:  true,
			Site: true,
			SiteConfig: &cra.Site{
				Templating: true,
				SSR:        true,
				Static:     true,
				SPA:        true,
			},
			Auth: cra.Session,
		},
	}

	if flagOptions.Debug {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			log.Printf("Failed to read build info")
			return
		}

		projectOptions.TemplateDir = `$GOPATH/pkg/mod/github.com/\!james\!tiberius\!kirk/go_web_template@` + bi.String()
	} else {
		projectOptions.TemplateDir = "."
	}

	bytes, _ := json.MarshalIndent(projectOptions, "", "    ")
	fmt.Println("Running with example project configuration")
	fmt.Println(string(bytes))

	tp := cra.NewWebTemplate(projectOptions, flagOptions.Verbose)
	err := tp.NewProject(projectOptions)
	if err != nil {
		log.Fatal(err)
	}
}