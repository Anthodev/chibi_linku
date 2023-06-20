package chibi_linku

import (
	"encoding/json"
	"net/http"
)

func getRootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		marshal, err := json.Marshal("Hello, World!")
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(marshal)
		if err != nil {
			return
		}
	}
}

func getEncodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleEncodeRequest(w, parseRequest(r.Body, w).Link)
	}
}
