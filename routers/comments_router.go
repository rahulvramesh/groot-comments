package routers

import (
	"github.com/gorilla/mux"
	"github.com/rahulvramesh/groot-comments/controllers"
	"github.com/urfave/negroni"
)

// SetCommentsRouter -
func SetCommentsRouter(router *mux.Router) *mux.Router {

	laRouter := mux.NewRouter()

	//find all data for spreadsheet/ csv
	laRouter.HandleFunc("/orgs/{orgName}/comments/", controllers.StoreCommentController).Methods("POST")

	router.PathPrefix("/orgs").Handler(negroni.New(
		negroni.Wrap(laRouter),
	))

	return router

}
