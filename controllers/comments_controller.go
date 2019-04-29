package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/evhivetech/pdfgo/common"
	"github.com/gorilla/mux"
	"github.com/rahulvramesh/groot-comments/models"

	jsoniter "github.com/json-iterator/go"
	validator "github.com/rahulvramesh/robo-validator"
	log "github.com/sirupsen/logrus"
)

//StoreCommentController - controller
func StoreCommentController(w http.ResponseWriter, r *http.Request) {

	//lets use the exisitng model than creating one
	var (
		payload *models.Comment
		err     error
	)

	//get the org name
	payload.AuthorID, _ = strconv.Atoi(r.Context().Value("author").(string))
	payload.Organization = mux.Vars(r)["orgName"]

	b, _ := ioutil.ReadAll(r.Body)
	err = jsoniter.Unmarshal(b, &payload)
	if err != nil {
		log.Error("Failed to unmarhsal json")
		common.DisplayAppError(w, errors.New("please check request payload data type"), "invalid filter request", http.StatusBadRequest)
		return
	}

	//my custom validator module for event based validation
	err = validator.Validate(payload, "store_comment")
	if err != nil {
		log.Error("Invalid payload")
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	//save to db

}
