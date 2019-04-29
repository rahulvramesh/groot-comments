package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/rahulvramesh/groot-comments/utils"

	log "github.com/sirupsen/logrus"
)

// FetchGrootUser -
func FetchGrootUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//Get the header and set in request as object
	userID, err := getXGrootUID(r)

	if err != nil {
		log.Error("Header Not Set")
		utils.DisplayAppError(w, err, "User id not present in request", http.StatusUnauthorized)

	} else {
		//else set the user in context
		//context.WithValue(r.Context(), "devconUser", userId)
		ctx := context.WithValue(r.Context(), "RequestBy", userID)

		//pass to next function
		next.ServeHTTP(w, r.WithContext(ctx))
	}

}

func getXGrootUID(r *http.Request) (string, error) {

	if userID := r.Header.Get("X-Groot-User"); userID != "" {
		return userID, nil
	}

	return "", errors.New("no user id present in the header")
}
