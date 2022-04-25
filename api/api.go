package api

import (
	"flag"
	"github.com/gorilla/mux"
	"go-todo-api/api/controllers"
	"go-todo-api/api/httpext"
	"go-todo-api/api/logger"
	"go-todo-api/api/routes"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "set api port")
	flag.Parse()
}

func Run() {
	defer logger.Flush()

	router := mux.NewRouter().StrictSlash(true)
	routes.WithLogger(router)

	controllers.RegisterTODOSControllers(router)

	log.Printf("API Listening http://localhost:%d\n\n", port)

	var (
		addr    = httpext.Port(port).Addr()
		handler = routes.WithCORS(router)
	)

	err := http.ListenAndServe(addr, routes.WithLogger(handler))
	if err != nil {
		logger.Error("error on server: %s", err.Error())
	}
}
