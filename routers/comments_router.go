package routers

import (
	"github.com/gorilla/mux"
	"github.com/rahulvramesh/groot-comments/controllers"
	"github.com/urfave/negroni"
)

// SetCommentsRouter -
func SetCommentsRouter(router *mux.Router) *mux.Router {

	commentRouter := mux.NewRouter()

	//find all data for spreadsheet/ csv
	commentRouter.HandleFunc("/orgs/{orgName}/comments/", controllers.StoreCommentController).Methods("POST")
	commentRouter.HandleFunc("/orgs/{orgName}/comments/", controllers.GetCommentsController).Methods("GET")
	commentRouter.HandleFunc("/orgs/{orgName}/comments/", controllers.DeleteCommentsController).Methods("DELETE")
	commentRouter.HandleFunc("/orgs/{orgName}/members/", controllers.GetAllMemberByOrgController).Methods("GET")

	router.PathPrefix("/orgs").Handler(negroni.New(
		negroni.Wrap(commentRouter),
	))

	return router

}
