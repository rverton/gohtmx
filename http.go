package gohtmx

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	addr   string
	server *http.Server
}

//go:embed templates/*
var templatesFS embed.FS

//go:embed public/*
var assetsFS embed.FS

var tplFuncs template.FuncMap

var (
	indexTmpl                = parseTpl(tplFuncs, "templates/index.html")
	selectTmpl               = parseTpl(tplFuncs, "templates/select.html")
	selectSpecializationTmpl = parseTpl(tplFuncs, "templates/partials/select-specialization.html")
)

type TplData struct {
	Data map[string]any
}

func NewServer(addr string) *Server {
	return &Server{
		server: &http.Server{
			Addr: addr,
		},
	}
}

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// display
	mux.HandleFunc("/", s.handleIndex)

	mux.HandleFunc("/select", s.handleSelect)
	mux.HandleFunc("/select-jobs", s.handleSelectJobs)

	mux.Handle("/public/", http.FileServer(http.FS(assetsFS)))

	return mux
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	indexTmpl.ExecuteTemplate(w, "layout.html", nil)
}

func (s *Server) handleSelect(w http.ResponseWriter, r *http.Request) {
	type Job struct {
		Name string
	}

	jobs := []Job{
		{Name: "Developer"},
		{Name: "Designer"},
		{Name: "Manager"},
	}

	selectTmpl.ExecuteTemplate(w, "layout.html", &TplData{
		Data: map[string]any{"Jobs": jobs},
	})
}

func (s *Server) handleSelectJobs(w http.ResponseWriter, r *http.Request) {

	job := r.URL.Query().Get("job")

	var specializations []string

	switch job {
	case "Developer":
		specializations = []string{"Frontend", "Backend", "Fullstack"}
	case "Designer":
		specializations = []string{"UI", "UX"}
	case "Manager":
		specializations = []string{"Project", "Team"}
	}

	selectSpecializationTmpl.ExecuteTemplate(w, "select-specialization.html", &TplData{
		Data: map[string]any{"Specializations": specializations},
	})
}

func (s *Server) Start() error {
	s.server.Handler = s.routes()

	log.Println("Starting server on", s.server.Addr)
	return s.server.ListenAndServe()
}

func parseTpl(funcs template.FuncMap, file string) *template.Template {

	tpls := []string{"templates/layout.html", "templates/partials/*.html", file}

	tpl, err :=
		template.New("layout.html").Funcs(funcs).ParseFS(
			templatesFS,
			tpls...,
		)
	if err != nil {
		log.Fatal("cant parse template", err)
	}

	return tpl
}
