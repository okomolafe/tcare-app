package httputil

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// RespondWithJSON is a helper method for writing JSON to an
// http.ResponseWriter. This helper will properly set the content-type before
// sending the JSON body.
func RespondWithJSON(w http.ResponseWriter, status int, resp interface{}) {
	json, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(json)))
	w.WriteHeader(status)
	w.Write(json)
}