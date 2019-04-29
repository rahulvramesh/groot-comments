package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rahulvramesh/groot-comments/models"
	"github.com/rahulvramesh/groot-comments/utils"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

//StoreCommentController - controller
func StoreCommentController(w http.ResponseWriter, r *http.Request) {

	//lets use the exisitng model than creating one
	var (
		payload *models.Comment
		err     error

		CommentModelObj models.CommentModel
	)

	//get the org name

	b, _ := ioutil.ReadAll(r.Body)
	err = jsoniter.Unmarshal(b, &payload)
	if err != nil {
		log.Error("Failed to unmarhsal json")
		utils.DisplayAppError(w, errors.New("please check request payload data type"), "invalid filter request", http.StatusBadRequest)
		return
	}

	payload.AuthorID, _ = strconv.Atoi(r.Context().Value("author").(string))
	payload.Organization = mux.Vars(r)["orgName"]

	//my custom validator module for event based validation
	// err = validator.Validate(payload, "store_comment")
	// if err != nil {
	// 	log.Error("Invalid payload")
	// 	utils.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	//save to db
	responseObj, err := CommentModelObj.StoreComment(payload)

	if err != nil {
		log.Error(err.Error())
		utils.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	//else print the response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	c, _ := jsoniter.Marshal(responseObj)
	w.Write(c)

}

//GetCommentsController -
func GetCommentsController(w http.ResponseWriter, r *http.Request) {

	var (
		CommentModelObj models.CommentModel
	)

	//get org name
	orgName := mux.Vars(r)["orgName"]

	//call the model
	responseObj, err := CommentModelObj.GetComments(orgName)

	if err != nil {
		log.Error(err.Error())
		utils.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	//else print the response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c, _ := jsoniter.Marshal(responseObj)
	w.Write(c)

}

//DeleteCommentsController -
func DeleteCommentsController(w http.ResponseWriter, r *http.Request) {

	var (
		CommentModelObj models.CommentModel
	)

	//get org name
	orgName := mux.Vars(r)["orgName"]

	//call delete model
	err := CommentModelObj.DeleteOrgComments(orgName)

	if err != nil {
		log.Error(err.Error())
		utils.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	//else print the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c, _ := jsoniter.Marshal(models.Comment{})
	w.Write(c)

}
