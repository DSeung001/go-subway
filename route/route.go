package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start(port int) {
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Printf("Listening on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
