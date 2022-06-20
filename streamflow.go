package streamflow

import (
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/collaborator"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/remote/coordinator"
	"github.com/zhenyiya/server/executor"
	"github.com/zhenyiya/server/mapper"
	"github.com/zhenyiya/server/reducer"
	"github.com/zhenyiya/server/task"
	"github.com/zhenyiya/server/workable"
	"github.com/zhenyiya/store"
	"github.com/zhenyiya/utils"
	"net/http"
	"os"
	"strconv"
)

func Set(key string, val ...interface{}) interface{} {
	cmd.Init()
	switch key {
	case constants.Mapper:
		fs := store.GetInstance()
		fs.SetMapper(val[0].(mapper.Mapper), val[1].(string))
	case constants.Reducer:
		fs := store.GetInstance()
		fs.SetReducer(val[0].(reducer.Reducer), val[1].(string))
	case constants.Executor:
		fs := store.GetInstance()
		fs.SetExecutor(val[0].(executor.Executor), val[1].(string))
	case constants.Function:
		// register function
		fs := store.GetInstance()
		f := val[0].(func(source *[]task.Countable,
			result *[]task.Countable,
			context *task.TaskContext) chan bool)
		if len(val) > 1 {
			fs.Add(f, val[1].(string))
			break
		}
		fs.Add(f)
	case constants.HashFunction:
		// register hash function
		fs := store.GetInstance()
		f := val[0].(func(source *[]task.Countable,
			result *[]task.Countable,
			context *task.TaskContext) chan bool)
		return fs.HAdd(f)
	case constants.Shared:
		fs := store.GetInstance()

		methods := val[0].([]string)
		handlers := make([]func(w http.ResponseWriter, r *http.Request) *task.Job, len(val)-1)
		for i, v := range val[1:] {
			handlers[i] = v.(func(w http.ResponseWriter, r *http.Request) *task.Job)
		}

		// register jobs
		fs.AddShared(methods, handlers...)
	case constants.Local:
		fs := store.GetInstance()

		methods := val[0].([]string)
		handlers := make([]func(w http.ResponseWriter, r *http.Request) *task.Job, len(val)-1)
		for i, v := range val[1:] {
			handlers[i] = v.(func(w http.ResponseWriter, r *http.Request) *task.Job)
		}

		// register jobs
		fs.AddLocal(methods, handlers...)
	case constants.ProjectPath:
		constants.ProjectDir = val[0].(string)
	}
	return nil
}

func Run(vars ...*cmd.SysVars) {
	// initialise environment
	cmd.Combine(append(vars, cmd.InitCmdEnv())...)

	var (
		localLogger *logger.Logger
		logFile     *os.File
	)

	runVars := cmd.Vars()

	switch runVars.CleanHistory {
	case constants.CleanHistory:
		localLogger, logFile = logger.NewLogger(runVars.LogPath, constants.DefaultLogPrefix, true)
	default:
		localLogger, logFile = logger.NewLogger(runVars.LogPath, constants.DefaultLogPrefix, true)
	}

	// set handler for router
	router := store.GetRouter()
	switch runVars.DebugMode {
	case true:
		router = utils.AdaptRouterToDebugMode(router)
	default:
	}

	logger.LogLogo("", "(c) 2017 zhenyiya", "", "Author: Hastings Yeung", "Github: https://github.com/HastingsYoung/zhenyiya", "")

	switch runVars.ServerMode {
	case constants.CollaboratorModeAbbr, constants.CollaboratorMode:
		// create collaborator
		clbt := collaborator.NewCollaborator()

		mst := workable.NewMaster()
		mst.BatchAttach(runVars.MaxRoutines)
		mst.LaunchAll()

		clbt.Join(mst)

		clbt.Handle(router)

	case constants.CoordinatorModeAbbr, constants.CoordinatorMode:
		cdnt := coordinator.GetCoordinatorInstance(runVars.Port)
		cdnt.Handle(router)
	}

	// launch server
	serv := &http.Server{
		Addr:        ":" + strconv.Itoa(runVars.Port),
		Handler:     router,
		ReadTimeout: constants.DefaultReadTimeout,
	}
	err := serv.ListenAndServe()
	logger.LogError(err.Error())
	localLogger.LogError(err.Error())
	logFile.Close()
}
