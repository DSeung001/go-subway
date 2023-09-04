package subway

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start(port int) {
	const apiPrefix = "/api/v1/"
	router := mux.NewRouter()

	router.Use(jsonContentTypeMiddleware)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	router.HandleFunc(apiPrefix+"getStationNames", handleGetStationName).Methods("GET")

	log.Printf("Listening on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// jsonContentTypeMiddleware : content type 을 json 으로 해더 설정
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// adapter 패턴
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			rw.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(rw, r)
	})
}

func handleGetStationName(rw http.ResponseWriter, r *http.Request) {
	var names = getStationNames()
	json.NewEncoder(rw).Encode(names)
}
