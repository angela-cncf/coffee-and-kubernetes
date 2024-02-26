package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (rest *RESTServer) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.RemoteAddr)
	message := rest.app.GetHelloMessage(r.RemoteAddr)
	json.NewEncoder(w).Encode(message)
}
