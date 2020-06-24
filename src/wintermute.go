package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func sendErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}

func index(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./static/index.html")
	if err != nil {
		sendErr(w, err)
		return
	}

	io.Copy(w, file)
}

func makeGenerator(chain frequencyChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		seed := r.FormValue("seed")
		seed = strings.TrimSpace(strings.ToLower(seed))
		if seed == "" {
			seed = "the"
		}

		io.WriteString(w, chain.generate(seed))
	}
}

func start() {
	chain, err := makeChain("./corpus")
	if err != nil {
		log.Fatalf("Could not instantiate chain, %s\n", err.Error())
	}

	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:7337",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.RequestURI, "/static") {
				log.Println(r.Method, r.RequestURI)
			}

			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/", index)
	r.Methods("GET").Path("/generate").Queries("seed", "{seed}").HandlerFunc(makeGenerator(chain))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Printf("Wintermute listening on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	start()
}
