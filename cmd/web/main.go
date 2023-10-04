package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zdev147/room_booking/pkg/config"
	"github.com/Zdev147/room_booking/pkg/handler"
	"github.com/Zdev147/room_booking/pkg/render"
)

const portNumber = ":8000"
const useCache = true

var appConfig config.AppConfig

func main() {
	initApp()

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(appConfig),
	}

	fmt.Printf("Server started at port %s\n", portNumber)
	server.ListenAndServe()
}

// initApp set all configuration required for app
func initApp() {
	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.TemplateCache = cache
	appConfig.UseCache = useCache

	render.SetConfig(&appConfig)

	repo := handler.CreateNewRepo(&appConfig)
	handler.CreateNewHanlder(&repo)
}
