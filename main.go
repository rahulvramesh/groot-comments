package main

import (
	"net/http"

	postgres "github.com/rahulvramesh/groot-comments/repository/postgres"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {

	//connect to db
	connectDB := postgres.New().Connect().GetSession()
	defer connectDB.Close()

	//rest of settings
	n := negroni.New()
	//set panic recovery
	n.Use(negroni.NewRecovery())

	server := &http.Server{
		Addr: "0.0.0.0:8005",
		// ReadTimeout:  time.Duration(svc.ReadTimeout) * time.Second,
		// WriteTimeout: time.Duration(svc.WriteTimeout) * time.Second,
		Handler: n,
	}

	log.WithFields(log.Fields{
		"request_id": "000",
	}).Info("Groot Comments Running At :8005")
	server.ListenAndServe()
}
