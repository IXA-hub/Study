package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
	"os"

	"github.com/matryer/goblueprints/chapter1/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", "8080", "The addr of the application.")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Webサーバーを開始します。port:", *addr)
	if err := http.ListenAndServe(":"+*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
