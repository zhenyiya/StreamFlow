package main

import (
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/core"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/remote/collaborator"
	"github.com/zhenyiya/remote/coordinator"
	"github.com/zhenyiya/server"
	"github.com/zhenyiya/server/workable"
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

	logger.LogHeader("Program Started")

	switch sysvars.ServerMode {
	case constants.CollaboratorModeAbbr, constants.CollaboratorMode:
		// create publisher
		pbls := server.GetPublisherInstance(localLogger)
		// create book keeper
		bkp := collaborator.NewBookKeeper(pbls, localLogger)
		bkp.NewBook()

		// register tasks
		pbls.AddShared(core.TaskAHandler, core.TaskBHandler, core.TaskCHandler)

		mst := workable.NewMaster(bkp, localLogger)
		mst.BatchAttach(sysvars.MaxRoutines)
		mst.LaunchAll()
		// connect to master
		pbls.Connect(mst)
		bkp.Watch(mst)

		bkp.Handle(router)
	case constants.CoordinatorModeAbbr, constants.CoordinatorMode:
		cdnt := coordinator.GetCoordinatorInstance(sysvars.Port, localLogger)
		cdnt.Handle(router)
	}

	// launch server
	serv := &http.Server{
		Addr:        ":" + strconv.Itoa(sysvars.Port),
		Handler:     router,
		ReadTimeout: constants.DefaultReadTimeout,
	}
	err := serv.ListenAndServe()
	logger.LogError(err.Error())
	localLogger.LogError(err.Error())
	logFile.Close()
}
