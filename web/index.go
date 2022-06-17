package web

import (
	"encoding/json"
	"github.com/zhenyiya/artifacts"
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/store"
	"github.com/zhenyiya/utils"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
)

type DataModel struct {
	Title   string
	Header1 string
	Header2 string
	Header3 string
	Body    []string
}

var templates = template.Must(template.ParseFiles(makePath("index.html")))
var validPath = regexp.MustCompile("^/(index)*$")

func Index(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "index")
	return
}

func Profile(w http.ResponseWriter, r *http.Request) {
	utils.AdaptHTTPWithHeader(w, constants.Header200OK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderContentTypeJSON)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, cmd.VarsJSONArrayStr())
}

func Routes(w http.ResponseWriter, r *http.Request) {
	router := store.GetRouter()

	base := artifacts.Base{"zhenyiya API", "[ Base URL: / ]"}
	entries := []artifacts.EntriesGroup{}
	models := []artifacts.ModelsGroup{}

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {

		es := []artifacts.Entry{}

		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		n := route.GetName()

		methods, err := route.GetMethods()
		if err != nil {
			return err
		}

		for _, m := range methods {
			es = append(es, artifacts.Entry{m, t, "", false})
		}

		entries = append(entries, artifacts.EntriesGroup{n, "", es})

		return nil
	})

	artifact := artifacts.Artifact{base, entries, models}
	mal, err := json.Marshal(artifact)

	if err != nil {
		panic(err)
	}

	utils.AdaptHTTPWithHeader(w, constants.Header200OK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderContentTypeJSON)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, string(mal))
}

func Logs(w http.ResponseWriter, r *http.Request) {
	str, err := logger.GetLogs()
	if err != nil {
		logger.LogError(err)
		return
	}
	utils.AdaptHTTPWithHeader(w, constants.Header200OK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, str)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makePath(path string) string {
	return filepath.Join(constants.LibUnixDir+"web", "templates", path)
}
