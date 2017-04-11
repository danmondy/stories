package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var templates *template.Template

func init() {
	ReadTemplates("views")
}

func main() {
	r := mux.NewRouter()
	serveStaticFiles(r)
	r.HandleFunc("/", App)
	http.ListenAndServe(":8081", r)

	log.Println("Fin.")
}

func App(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{}, 0)
	renderTemplate(w, "app", m)
}

func GetRandomCharacter() *Character {

	now := time.Now()
	rand.Seed(now.UnixNano())

	return &Character{}
}

func serveStaticFiles(r *mux.Router) {
	fs := http.FileServer(http.Dir("views"))
	r.PathPrefix("/img").Handler(fs)
	r.PathPrefix("/css").Handler(fs)
	r.PathPrefix("/components").Handler(fs)
	r.PathPrefix("/fonts").Handler(fs)
}

func ReadTemplates(templatePath string) {
	fmt.Println(templatePath)

	templates = template.New("whatisthis").Funcs(template.FuncMap{})
	err := filepath.Walk(templatePath,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println("template:", strings.TrimPrefix(strings.Replace(path, "\\", "/", -1), "views/"))
			if strings.Contains(path, ".html") {
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}
				fmt.Printf("views%q", os.PathSeparator)
				_, err = templates.New(strings.TrimPrefix(strings.Replace(path, "\\", "/", -1), "views/")).Parse(string(bytes))
				if err != nil {
					panic(err)
				}
			}
			return err
		})
	if err != nil {
		panic(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, model map[string]interface{}) error {

	if model == nil {
		model = make(map[string]interface{})
	}
	model["title"] = tmpl
	fmt.Println(model["title"])
	err := templates.ExecuteTemplate(w, tmpl+".html", model)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
