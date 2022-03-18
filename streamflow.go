package streamflow

import (
	"github.com/zhenyiya/cmd"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/funcstore"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/remote/collaborator"
	"github.com/zhenyiya/remote/coordinator"
	"github.com/zhenyiya/server"
	"github.com/zhenyiya/server/mapper"
	"github.com/zhenyiya/server/reducer"
	"github.com/zhenyiya/server/task"
	"github.com/zhenyiya/server/workable"
	"github.com/zhenyiya/utils"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var singleton *Vars
var once sync.Once

func Set(key string, val ...interface{}) {
	Init()
	switch key {
	case constants.Mapper:
		singleton.Mapper = val[0].(mapper.Mapper)
	case constants.Reducer:
		singleton.Reducer = val[0].(reducer.Reducer)
	case constants.Function:
		// register function
		fs := funcstore.GetFSInstance()
		f := val[0].(func(source *[]task.Countable,
			result *[]task.Countable,
			context *task.TaskContext) chan bool)
		if len(val) > 1 {
			fs.Add(f, val[1].(string))
			break
		}
		fs.Add(f)
	case constants.Shared:
		pbls := server.GetPublisherInstance()

		methods := val[0].([]string)
		handlers := make([]func(w http.ResponseWriter, r *http.Request) task.Task, len(val)-1)
		for i, v := range val[1:] {
			handlers[i] = v.(func(w http.ResponseWriter, r *http.Request) task.Task)
		}

		// register tasks
		pbls.AddShared(methods, handlers...)
	case constants.Local:
		pbls := server.GetPublisherInstance()

		methods := val[0].([]string)
		handlers := make([]func(w http.ResponseWriter, r *http.Request) task.Task, len(val)-1)
		for i, v := range val[1:] {
			handlers[i] = v.(func(w http.ResponseWriter, r *http.Request) task.Task)
		}

		// register tasks
		pbls.AddLocal(methods, handlers...)
	case constants.ProjectPath:
		constants.ProjectDir = val[0].(string)
	}
}

func Run(vars ...*Vars) {
	Init()
	// initialise environment
	sysvars := cmd.InitCmdEnv()
	singleton = combine(sysvars, vars...)

	var (
		localLogger *logger.Logger
		logFile     *os.File
	)

	switch singleton.CleanHistory {
	case constants.CleanHistory:
		localLogger, logFile = logger.NewLogger(singleton.LogPath, constants.DefaultLogPrefix, true)
	default:
		localLogger, logFile = logger.NewLogger(singleton.LogPath, constants.DefaultLogPrefix, true)
	}

	// set handler for router
	router := mux.NewRouter()
	switch singleton.DebugMode {
	case true:
		router = utils.AdaptRouterToDebugMode(router)
	default:
	}

	logger.LogHeader("Program Started")

	switch singleton.ServerMode {
	case constants.CollaboratorModeAbbr, constants.CollaboratorMode:
		// create publisher
		pbls := server.GetPublisherInstance()
		server.Logger(localLogger)
		// create book keeper
		bkp := collaborator.NewBookKeeper(pbls, localLogger)
		bkp.NewBook()

		mst := workable.NewMaster(bkp, localLogger)

		mst.Mapper(singleton.Mapper).Reducer(singleton.Reducer)

		mst.BatchAttach(singleton.MaxRoutines)
		mst.LaunchAll()

		// connect to master
		pbls.Connect(mst)
		bkp.Watch(mst)

		bkp.Handle(router)

	case constants.CoordinatorModeAbbr, constants.CoordinatorMode:
		cdnt := coordinator.GetCoordinatorInstance(singleton.Port, localLogger)
		cdnt.Handle(router)
	}

	// launch server
	serv := &http.Server{
		Addr:        ":" + strconv.Itoa(singleton.Port),
		Handler:     router,
		ReadTimeout: constants.DefaultReadTimeout,
	}
	err := serv.ListenAndServe()
	logger.LogError(err.Error())
	localLogger.LogError(err.Error())
	logFile.Close()
}

type Vars struct {
	ServerMode      string
	DebugMode       bool
	Port            int
	ContactBookPath string
	LogPath         string
	DataStorePath   string
	MaxRoutines     int
	CleanHistory    bool
	Mapper          mapper.Mapper
	Reducer         reducer.Reducer
}

func Init() {
	once.Do(func() {
		singleton = &Vars{}
		path, err := filepath.Abs("./")
		if err != nil {
			panic(err)
		}
		constants.ProjectDir = path
	})
}

func combine(sysVars *cmd.SysVars, vars ...*Vars) *Vars {
	var v Vars
	if len(vars) > 0 {
		v = *vars[0]
	} else {
		v = *singleton
	}
	s := *sysVars
	v.ServerMode = s.ServerMode
	v.DebugMode = s.DebugMode
	v.Port = s.Port
	v.ContactBookPath = s.ContactBookPath
	v.LogPath = s.LogPath
	v.DataStorePath = s.DataStorePath
	v.MaxRoutines = s.MaxRoutines
	v.CleanHistory = s.CleanHistory
	return &v
}
