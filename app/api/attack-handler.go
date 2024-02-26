package api

import (
	"fmt"
	"net/http"
	"os"
)

func (rest *RESTServer) handleAttack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming attack from " + r.RemoteAddr)
	os.Exit(13)
}
