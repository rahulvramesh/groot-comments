package main

import (
	"net/http"

	"github.com/rahulvramesh/groot-comments/db"

	"github.com/rahulvramesh/groot-comments/middlewares"
	"github.com/rahulvramesh/groot-comments/routers"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"
)

func main() {

	//init db
	db.Connect()
	defer db.GetSession().Close()

	//initialize router
	router := routers.InitRoutes()

	//migaration
	//db.GetSession().AutoMigrate(&models.Comment{})

	//set middlewere for user id
	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.FetchGrootUser))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8005",
		Handler: n,
	}

	log.Info("Listening on port!!!")
	server.ListenAndServe()
}
