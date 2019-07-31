package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// Setup our Ctrl+C handler
	SetupCloseHandler()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/parse", handleParse)

	log.Println("Listening on localhost: " + getPort())
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleParse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		inputFormat := r.FormValue("inputFormat")
		outputFormat := r.FormValue("outputFormat")
		inputContent := r.FormValue("inputContent")

		fmt.Println(inputFormat + " -> " + outputFormat)
		fmt.Println(inputContent)
	}
	http.Redirect(w, r, "/", 302)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	dat := struct {
		Year int
	}{time.Now().Year()}

	tpl.ExecuteTemplate(w, "index.gohtml", dat)
}

// TODO: Decide on a port or search for an open one.
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
// Thanks Edd: https://golangcode.com/handle-ctrl-c-exit-in-terminal/
func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rCtrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
