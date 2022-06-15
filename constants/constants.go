package constants

import (
	"errors"
	"os"
	"time"
)

// system vars
const (
	CollaboratorModeAbbr   = "clbt"
	CollaboratorMode       = "collaborator"
	CoordinatorModeAbbr    = "cdnt"
	CoordinatorMode        = "coordinator"
	DebugInactivated       = false
	DebugActivated         = true
	DefaultListenPort      = 8080
	DefaultCasePath        = "./case.json"
	DefaultLogPath         = "./history.log"
	DefaultDataStorePath   = "./streamflow.dat"
	DefaultLogPrefix       = "zhenyiya:"
	CleanHistory           = true
	DefaultNotCleanHistory = false
	Mapper                 = "Mapper"
	Reducer                = "Reducer"
	Function               = "Function"
	HashFunction           = "HashFunction"
	Shared                 = "Shared"
	Local                  = "Local"
	ProjectPath            = "ProjectPath"
	DefaultWorkerPerMaster = 10
	DefaultHost            = "localhost"
	DefaultGossipNum       = 5
	DefaultCaseID          = "zhenyiyaStandardCase"
)

// store setting
const (
	DefaultHashLength = 12
)

var (
	DefaultReadTimeout         = 15 * time.Second
	DefaultPeriodShort         = 500 * time.Millisecond
	DefaultPeriodLong          = 2000 * time.Millisecond
	DefaultPeriodRoutineDay    = 24 * time.Hour
	DefaultPeriodRoutineWeek   = 7 * 24 * time.Hour
	DefaultPeriodRoutine30Days = 30 * DefaultPeriodRoutineDay
	DefaultPeriodPermanent     = 0 * time.Second
	DefaultTaskExpireTime      = 30 * time.Second
	DefaultGCInterval          = 30 * time.Second
	DefaultMaxMappingTime      = 600 * time.Second
	DefaultSynInterval         = 3 * time.Minute
)

// executor types
const (
	ExecutorTypeMapper   = "mapper"
	ExecutorTypeReducer  = "reducer"
	ExecutorTypeShuffler = "shuffler"
	ExecutorTypeCombiner = "combiner"
	ExecutorTypeDefault  = "default"
)

// communication types
const (
	ArgTypeInteger              = "integer"
	ArgTypeNumber               = "number"
	ArgTypeString               = "string"
	ArgTypeObject               = "object"
	ArgTypeBoolean              = "boolean"
	ArgTypeNull                 = "null"
	ArgTypeArray                = "array"
	ConstraintTypeMax           = "maximum"
	ConstraintTypeMin           = "minimum"
	ConstraintTypeXMin          = "exclusiveMinimum"
	ConstraintTypeXMax          = "exclusiveMaximum"
	ConstraintTypeUniqueItems   = "uniqueItems"
	ConstraintTypeMaxProperties = "maxProperties"
	ConstraintTypeMinProperties = "minProperties"
	ConstraintTypeMaxLength     = "maxLength"
	ConstraintTypeMinLength     = "minLength"
	ConstraintTypePattern       = "pattern"
	ConstraintTypeMaxItems      = "maxItems"
	ConstraintTypeMinItems      = "minItems"
	ConstraintTypeEnum          = "enum" // value of interface{} should accept a slice
	ConstraintTypeAllOf         = "allOf"
	ConstraintTypeAnyOf         = "anyOf"
	ConstraintTypeOneOf         = "oneOf"
)

// errors
var (
	ErrUnknownCmdArg                   = errors.New("zhenyiya: unknown commandline argument, please enter -h to check out")
	ErrConnectionClosed                = errors.New("zhenyiya: connection closed")
	ErrUnknown                         = errors.New("zhenyiya: unknown error")
	ErrAPIError                        = errors.New("zhenyiya: api error")
	ErrNoCollaborator                  = errors.New("zhenyiya: collaborator does not exist")
	ErrCollaboratorExists              = errors.New("zhenyiya: collaborator already exists")
	ErrNoService                       = errors.New("zhenyiya: service of id does not exist")
	ErrConflictService                 = errors.New("zhenyiya: found conflict, service of id already exists")
	ErrNoRegister                      = errors.New("zhenyiya: register does not exist")
	ErrConflictRegister                = errors.New("zhenyiya: found conflict, provider of the service already exists")
	ErrNoSubscriber                    = errors.New("zhenyiya: subscriber does not exist")
	ErrConflictSubscriber              = errors.New("zhenyiya: found conflict, subscriber of the service already exists")
	ErrTimeout                         = errors.New("zhenyiya: task timeout error")
	ErrNoPeers                         = errors.New("zhenyiya: no peer appears in the contact book")
	ErrFunctNotExist                   = errors.New("zhenyiya: no such function found in store")
	ErrJobNotExist                     = errors.New("zhenyiya: no sucn job found in store")
	ErrExecutorNotFound                = errors.New("zhenyiya: no such executor found in store")
	ErrValNotFound                     = errors.New("zhenyiya: no value found with such key")
	ErrCaseMismatch                    = errors.New("zhenyiya: case mismatch error")
	ErrCollaboratorMismatch            = errors.New("zhenyiya: collaborator mismatch error")
	ErrMapTaskFailing                  = errors.New("zhenyiya: map operation failing error")
	ErrReduceTaskFailing               = errors.New("zhenyiya: reduce operation failing error")
	ErrExecutorStackLengthInconsistent = errors.New("zhenyiya: executor stack length inconsistent error")
)

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HTTP headers
var (
	Header200OK        = Header{"200", "OK"}
	Header201Created   = Header{"201", "Created"}
	Header202Accepted  = Header{"202", "Accepted"}
	Header204NoContent = Header{"204", "NoContent"}
	Header403Forbidden = Header{"403", "Forbidden"}
	Header404NotFound  = Header{"404", "NotFound"}
	Header409Conflict  = Header{"409", "Conflict"}
)

// Gossip Protocol headers
var (
	GossipHeaderOK                   = Header{"200", "OK"}
	GossipHeaderCaseMismatch         = Header{"401", "CaseMismatch"}
	GossipHeaderCollaboratorMismatch = Header{"401", "CollaboratorMismatch"}
	GossipHeaderUnknownError         = Header{"500", "UnknownGossipError"}
)

// restful
const (
	JSONAPIVersion = `{"version":"1.0"}`
)

var (
	ProjectDir     = ""
	ProjectUnixDir = ""
	LibDir         = "github.com/zhenyiya/"
	LibUnixDir     = os.Getenv("GOPATH") + "/src/github.com/zhenyiya/"
)
