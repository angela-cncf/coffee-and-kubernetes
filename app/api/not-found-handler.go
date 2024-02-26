package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type notFoundHandler struct {
}

func (th *notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": fmt.Sprintf("I don't understand what you are asking %s. Try hello instead.", r.RemoteAddr)})
}
