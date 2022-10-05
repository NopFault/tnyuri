package tnyuri

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitWeb() {

	// set routes
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/{short:[a-zA-Z0-9]+}", followShort)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Server is down: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("It's Working"))
}

func followShort(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	short := params["short"]

	var item []URL = By[URL]("short", short)

	if len(item) > 0 {
		// Update stats
		stats := item[0].Stats()
		stats.Increase()

		// Redirect to url
		http.Redirect(w, r, item[0].Url, 301)
	}
}
