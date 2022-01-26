package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"restfulAPI/httputil"
	"restfulAPI/users"
)


func UserPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	up, err := users.GetUserPosts(id)
	if err != nil {
		httputil.RespondWithJSON(w, 500, err)
		return
	}

	if up.UserId == 0 {
		httputil.RespondWithJSON(w, 404, struct{}{})
		return
	}

	httputil.RespondWithJSON(w, 200, up)
}