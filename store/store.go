package store

import (
	"github.com/zhenyiya/StreamFlow/artifacts/iexecutor"
	"github.com/zhenyiya/StreamFlow/artifacts/imapper"
	"github.com/zhenyiya/StreamFlow/artifacts/ireducer"
	"github.com/zhenyiya/StreamFlow/artifacts/message"
	"github.com/zhenyiya/StreamFlow/artifacts/task"
	"github.com/zhenyiya/StreamFlow/constants"
	"github.com/zhenyiya/StreamFlow/logger"
	"github.com/zhenyiya/StreamFlow/utils"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

var router *mux.Router
var msgChan chan *message.CardMessageFuture
var singleton *FS
var once sync.Once
var onceRouter sync.Once
var onceMsgChan sync.Once
var mu sync.Mutex

type color int

const (
	white color = iota
	grey
	black
)

func GetRouter() *mux.Router {
	onceRouter.Do(func() {
		router = mux.NewRouter()
	})
	return router
}

func GetMsgChan() chan *message.CardMessageFuture {
	onceMsgChan.Do(func() {
		msgChan = make(chan *message.CardMessageFuture)
	})
	return msgChan
}

func GetInstance() *FS {
	once.Do(func() {
		singleton = &FS{make(map[string]func(source *task.Collection,
			result *task.Collection,
			context *task.TaskContext) bool),
			make(map[string]*color),
			make(map[string]iexecutor.IExecutor),
			make(map[string]*task.Job),
			make(map[string]*JobFunc),
			make(map[string]*JobFunc),
			make(map[string]*rate.Limiter)}
		singleton.sweep()
	})
	return singleton
}

type FS struct {
	Funcs map[string]func(source *task.Collection,
		result *task.Collection,
		context *task.TaskContext) bool
	memstack   map[string]*color
	executors  map[string]iexecutor.IExecutor
	jobs       map[string]*task.Job
	SharedJobs map[string]*JobFunc
	LocalJobs  map[string]*JobFunc
	limiters   map[string]*rate.Limiter
}

type JobFunc struct {
	F         func(w http.ResponseWriter, r *http.Request, bg *task.Background)
	Methods   []string
	Signature string
}

func (fs *FS) Add(f func(source *task.Collection,
	result *task.Collection,
	context *task.TaskContext) bool, id ...string) {
	var i string
	if len(id) < 1 {
		i = utils.StripRouteToFunctName(utils.ReflectFuncName(f))
	} else {
		i = id[0]
	}

	mu.Lock()
	defer mu.Unlock()
	fs.Funcs[i] = f
}

func (fs *FS) HAdd(f func(source *task.Collection,
	result *task.Collection,
	context *task.TaskContext) bool) (hash string) {
	hash = utils.RandStringBytesMaskImprSrc(constants.DEFAULT_HASH_LENGTH)

	mu.Lock()
	defer mu.Unlock()
	fs.Funcs[hash] = f
	// fs.Outbound[hash] = make(chan bool)
	*fs.memstack[hash] = grey
	return
}

func (fs *FS) Call(id string, source *task.Collection,
	result *task.Collection,
	context *task.TaskContext) bool {

	var (
		bol bool = false
	)

	if f := fs.Funcs[id]; f != nil {
		if c := fs.memstack[id]; c != nil {
			bol = f(source, result, context)
			*fs.memstack[id] = white
			return bol
		}
		bol = f(source, result, context)
		return bol
	}

	logger.LogError(constants.ERR_FUNCT_NOT_EXIST)
	return bol
}

func (fs *FS) SetMapper(mp imapper.IMapper, name string) {
	exe := iexecutor.Default()
	exe.Todo(mp.Map)
	exe.Type(constants.EXECUTOR_TYPE_MAPPER)
	fs.executors[name] = exe
}

func (fs *FS) SetReducer(rd ireducer.IReducer, name string) {
	exe := iexecutor.Default()
	exe.Todo(rd.Reduce)
	exe.Type(constants.EXECUTOR_TYPE_REDUCER)
	fs.executors[name] = exe
}

func (fs *FS) SetExecutor(exe iexecutor.IExecutor, name string) {
	fs.executors[name] = exe
}

func (fs *FS) GetExecutor(name string) (iexecutor.IExecutor, error) {
	if exe := fs.executors[name]; exe != nil {
		return exe, nil
	}
	return iexecutor.Default(), constants.ERR_EXECUTOR_NOT_FOUND
}

func (fs *FS) SetJob(j *task.Job) {
	fs.jobs[j.Id()] = j
}

func (fs *FS) GetJob(id string) (*task.Job, error) {
	if j := fs.jobs[id]; j != nil {
		return j, nil
	}
	return task.MakeJob(), constants.ERR_JOB_NOT_EXIST
}

func (fs *FS) SetShared(key string, val *JobFunc) {
	fs.SharedJobs[key] = val
}

func (fs *FS) SetLocal(key string, val *JobFunc) {
	fs.LocalJobs[key] = val
}

func (fs *FS) GetLocal(key string) (*JobFunc, error) {
	if j := fs.LocalJobs[key]; j != nil {
		return j, nil
	}
	return new(JobFunc), constants.ERR_JOB_NOT_EXIST
}

func (fs *FS) GetShared(key string) (*JobFunc, error) {
	if j := fs.SharedJobs[key]; j != nil {
		return j, nil
	}
	return new(JobFunc), constants.ERR_JOB_NOT_EXIST
}

func (fs *FS) AddLocal(methods []string, jobs ...func(w http.ResponseWriter, r *http.Request, bg *task.Background)) {
	for _, f := range jobs {
		signature := utils.StripRouteToAPIRoute(utils.ReflectFuncName(f))
		fs.LocalJobs[signature] = &JobFunc{f, methods, signature}
	}
}

func (fs *FS) AddShared(methods []string, jobs ...func(w http.ResponseWriter, r *http.Request, bg *task.Background)) {
	for _, f := range jobs {
		signature := utils.StripRouteToAPIRoute(utils.ReflectFuncName(f))
		fs.SharedJobs[signature] = &JobFunc{f, methods, signature}
	}
}

func (fs *FS) SetLimiter(name string, r rate.Limit, b int) {
	fs.limiters[name] = rate.NewLimiter(r, b)
}

func (fs *FS) GetLimiter(name string) (*rate.Limiter, error) {
	lim := fs.limiters[name]
	if lim != nil {
		return lim, nil
	}
	return lim, constants.ERR_LIMITER_NOT_FOUND
}

func (fs *FS) sweep() {
	go func() {
		for {
			<-time.After(constants.DEFAULT_GC_INTERVAL)
			// copy lookup table
			stack := fs.memstack
			for k, s := range stack {
				if *s == white {
					fs.delete(k)
				}
			}
		}
	}()
}

func (fs *FS) delete(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(fs.Funcs, id)
	// delete(fs.Outbound, id)
	delete(fs.memstack, id)
}
