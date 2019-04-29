// router index file, initialize all router here (files)
package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	router = SetCommentsRouter(router)

	return router
}
