package web

import (
	"encoding/json"
	"github.com/zhenyiya/artifacts/restful"
	"github.com/zhenyiya/artifacts/stats"
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/store"
	"github.com/zhenyiya/utils"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	utils.AdaptHTTPWithStatus(w, http.StatusOK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderContentTypeJSON)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, cmd.VarsJSONArrayStr())
}

func Routes(w http.ResponseWriter, r *http.Request) {
	router := store.GetRouter()

	base := restful.Base{"zhenyiya API", "[ Base URL: / ]"}
	entries := []restful.EntriesGroup{}
	models := []restful.ModelsGroup{}

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {

		es := []restful.Entry{}

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
			es = append(es, restful.Entry{m, t, "", false})
		}

		entries = append(entries, restful.EntriesGroup{n, "", es})

		return nil
	})

	dbPayload := restful.DashboardPayload{base, entries, models}
	mal, err := json.Marshal(dbPayload)

	if err != nil {
		panic(err)
	}

	utils.AdaptHTTPWithStatus(w, http.StatusOK)
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
	utils.AdaptHTTPWithStatus(w, http.StatusOK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, str)
}

func Stats(w http.ResponseWriter, r *http.Request) {
	mal, err := json.Marshal(stats.GetStatsInstance().Stats())

	if err != nil {
		panic(err)
	}

	utils.AdaptHTTPWithStatus(w, http.StatusOK)
	utils.AdaptHTTPWithHeader(w, constants.HeaderContentTypeJSON)
	utils.AdaptHTTPWithHeader(w, constants.HeaderCORSEnableAllOrigin)
	io.WriteString(w, string(mal))
}
