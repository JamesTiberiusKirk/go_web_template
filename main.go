package main

import (
	"encoding/json"
	"fmt"
	"music_manager/session"
	"music_manager/site"
)

func main() {
	initLogger()

	config := buildConfig()

	db := initDb(config)
	e := initServer(config)

	sessionManager := session.New()

	s := site.NewSite(config.Http.RootPath, db, sessionManager, e, config.SessionSecret)
	s.Serve()

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	fmt.Print(string(data))

	e.Logger.Fatal(e.Start(config.Http.Port))
}