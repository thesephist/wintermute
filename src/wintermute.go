package main

import (
	"html/template"
	"io"
	"io/ioutil"
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

		io.WriteString(w, chain.generateFrom(seed, 150))
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type FullPageTemplateVars struct {
	Title        string
	Date         time.Time
	ContentLines []string
	NextTitle    string
}

func makeFullPager(chain frequencyChain) http.HandlerFunc {
	file, err := os.Open("./templates/full.html")
	check(err)
	defer file.Close()

	part, err := ioutil.ReadAll(file)
	check(err)

	tmpl, err := template.New("full").Parse(string(part))
	check(err)

	return func(w http.ResponseWriter, r *http.Request) {
		title := strings.ReplaceAll(chain.generate(3), "\n", " ")
		nextTitle := strings.ReplaceAll(chain.generate(7), "\n", " ")
		contentLines := strings.Split(chain.generate(800), "\n")

		vars := FullPageTemplateVars{
			Title:        title,
			Date:         time.Now(),
			ContentLines: contentLines,
			NextTitle:    nextTitle,
		}

		err := tmpl.Execute(w, vars)
		if err != nil {
			sendErr(w, err)
		}
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
	r.Methods("GET").Path("/full").HandlerFunc(makeFullPager(chain))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Printf("Wintermute listening on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	start()
}
