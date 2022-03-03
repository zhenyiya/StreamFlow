package main

import (
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/core"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/remote"
	"github.com/zhenyiya/server"
	"github.com/zhenyiya/utils"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// initialise environment
	sysvars := cmd.InitCmdEnv()
	var (
		localLogger *logger.Logger
		logFile     *os.File
	)

	switch sysvars.CleanHistory {
	case constants.CleanHistory:
		localLogger, logFile = logger.NewLogger(sysvars.LogPath, constants.DefaultLogPrefix, true)
	default:
		localLogger, logFile = logger.NewLogger(sysvars.LogPath, constants.DefaultLogPrefix, true)
	}
	// set handler for router
	router := mux.NewRouter()
	switch sysvars.DebugMode {
	case true:
		router = utils.AdaptRouterToDebugMode(router)
	default:
	}

	switch sysvars.ServerMode {
	case constants.CollaboratorModeAbbr, constants.CollaboratorMode:
		// create book keeper
		bkp := remote.NewBookKeeper()
		// create publisher
		pbls := server.GetPublisherInstance(localLogger)
		bkp.WatchNewBook(pbls, localLogger)

		// register tasks
		pbls.AddExposed(core.TaskAHandler, core.TaskBHandler, core.TaskCHandler)

		mst := server.NewMaster(localLogger)
		mst.BatchAttach(sysvars.MaxRoutines)
		mst.LaunchAll()
		// connect to master
		pbls.Connect(mst)
		bkp.Handle(router)
	case constants.CoordinatorModeAbbr, constants.CoordinatorMode:
		regCenter := remote.GetRegCenterInstance(sysvars.Port, localLogger)
		regCenter.Handle(router)
	}

	// launch server
	serv := &http.Server{
		Addr:        ":" + strconv.Itoa(sysvars.Port),
		Handler:     router,
		ReadTimeout: constants.DefaultReadTimeout,
	}
	err := serv.ListenAndServe()
	localLogger.LogError(err.Error())
	logFile.Close()
}
